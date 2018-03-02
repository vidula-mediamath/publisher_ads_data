package main

import (
	"errors"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"log"
	//"fmt"
	"sync"
	//"golang.org/x/text/encoding"
	//"golang.org/x/text/transform"
	"gopkg.in/iconv.v1"
)

func main() {
	db, err := storage.NewPostgres()
	defer db.Close()
	if err != nil{
		return
	}
	var wg sync.WaitGroup
	urls := []string{"https://www.yahoo.com/ads.txt", "https://www.cnn.com/ads.txt", "http://www.nytimes.com/ads.txt", "https://www.nbc.com/ads.txt"}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL
			pubName, err := GetPublisherName(url)
			if err != nil {
				return
			}
			body, err := ExecuteGetOnAdsPage(url)
			if err != nil {
				return
			}
			records, err := ParseHttpResp(body)
			if err != nil {
				return
			}
			dbInsertErrors := db.DBInsert(records, pubName)
			log.Println(dbInsertErrors)
			log.Println(err)
		}(url)
	}
	wg.Wait()
}

//GetPublisherName retrieves the publisher domain from the urls so that later it can be used to store in the db
func GetPublisherName(pubUrl string) (string, error) {
	u, err := url.Parse(pubUrl)
	if err != nil {
		return "", err
	}
	return u.Hostname(), err
}

func ExecuteGetOnAdsPage(pubUrl string) ([]byte, error) {
	resp, err := http.Get(pubUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	if resp.Header.Get("Content-Type") != "" {
		if !strings.Contains(resp.Header.Get("Content-Type"), "utf-8") {
			receivedCharEncoding := strings.Split(resp.Header.Get("Content-Type"), ";")[1]
			if !strings.Contains(receivedCharEncoding, "utf-8"){
				usedEncoding := strings.Split(receivedCharEncoding, "charset=")[1]
				cd, err := iconv.Open("utf-8", usedEncoding) // convert usedEncoding to utf8
				if err != nil {
					return nil, err
				}
				defer cd.Close()
				
				var converted []byte = make([]byte, 10000000)
				decodedBody, _, err := cd.Conv(body, converted)
				return decodedBody, err			
			}
		}
	}	
	return body, err
}

func ParseHttpResp(body []byte) ([]storage.Record, error) {
	var FileData []storage.Record
	var s1 []string = strings.Split(string(body), "\n")

	for _, v := range s1 {
		var record storage.Record
		var splitBeforeComment []string = strings.Split(v, "#")

		var splitOnEachComma []string = strings.Split(splitBeforeComment[0], ",")
		if len(splitOnEachComma) >= 3 {
			domain, err := validateSupplyDomain(splitOnEachComma[0])
			if err != nil {
				continue
			}
			relation, err := validateRelationValue(splitOnEachComma[2])
			if err != nil {
				continue
			}
			record.Supply_source_domain = domain
			record.Id = strings.TrimSpace(splitOnEachComma[1])
			record.Relationship = relation

			FileData = append(FileData, record)
		}
	}
	return FileData, nil
}

func validateSupplyDomain(input string) (string, error) {
	input = strings.TrimSpace(input)
	if !strings.Contains(input, ".") {
		return "", errors.New("Invalid supply domain name")
	}
	return input, nil
}

func validateRelationValue(input string) (string, error) {
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	switch input {
	case "DIRECT":
		return input, nil
	case "RESELLER":
		return input, nil
	default:
		return "", errors.New("Invalid relation value")
	}
}
