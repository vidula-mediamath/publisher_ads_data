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
	"regexp"
)

func main() {
	var wg sync.WaitGroup
	r, _ := regexp.Compile("^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\\.([a-zA-Z]{2,}|[a-zA-Z0-9-]{2,30}\\.[a-zA-Z]{2,3})$")
	urls := []string{"https://www.yahoo.com/ads.txt", "https://www.cnn.com/ads.txt", "http://www.nytimes.com/ads.txt", "https://www.nbc.com/ads.txt"}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL
			pubName, err := getPublisherName(url)
			if err != nil {
				log.Println(err)
				return
			}
			body, err := executeGetOnAdsPage(url)
			if err != nil {
				log.Println(err)
				return
			}
			records, err := parseHttpResp(body, r)
			if err != nil {
				log.Println(err)
				return
			}
			dbErr := storage.AddRecordsInDB(records, pubName)
			if dbErr != nil {
				log.Println(dbErr)
				return
			}
		}(url)
	}
	wg.Wait()
}

//GetPublisherName retrieves the publisher domain from the urls so that later it can be used to store in the db
func getPublisherName(pubUrl string) (string, error) {
	u, err := url.Parse(pubUrl)
	if err != nil {
		return "", err
	}
	return u.Hostname(), err
}

func executeGetOnAdsPage(pubUrl string) ([]byte, error) {
	resp, err := http.Get(pubUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func parseHttpResp(body []byte, r *Regexp) ([]storage.Record, error) {
	var FileData []storage.Record
	var lines []string = strings.Split(string(body), "\n")

	for _, line := range lines {
		var record storage.Record
		var splitBeforeComment []string = strings.Split(line, "#")

		var splitOnEachComma []string = strings.Split(splitBeforeComment[0], ",")
		
		//validate record and its data
		if len(splitOnEachComma) >= 3 {
			
			//transform the data, TrimSpace and ToUpper
			for i:=0; i<len(splitOnEachComma); i++{
				splitOnEachComma[i] = strings.TrimSpace(splitOnEachComma[i])
			}
			splitOnEachComma[2] = strings.ToUpper(splitOnEachComma[2])
			
			//validate suuply domain value
			if !r.MatchString(splitOnEachComma[0]) {
				return errors.New("Invalid supply domain name")
			}
			//validate relation value
			err := validateRelationValue(splitOnEachComma[2])
			if err != nil {
				continue
			}
			
			//assign the parsed input values to record			
			record.Supply_source_domain = splitOnEachComma[0]
			record.Id = splitOnEachComma[1]
			record.Relationship = strings.ToUpper(splitOnEachComma[2])

			FileData = append(FileData, record)
		}
	}
	return FileData, nil
}

//func validateSupplyDomain(input string,) error {
//	r, _ := regexp.Compile("^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\\.([a-zA-Z]{2,}|[a-zA-Z0-9-]{2,30}\\.[a-zA-Z]{2,3})$")	
//	if !r.MatchString(input) {
//		return errors.New("Invalid supply domain name")
//	}
//	return nil
//}

func validateRelationValue(input string) error {
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	switch input {
	case "DIRECT":
		return nil
	case "RESELLER":
		return nil
	default:
		return errors.New("Invalid relation value")
	}
}