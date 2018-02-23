package main

import (
	"errors"
	"fmt"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func main() {
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
				log.Println("Invalid url")
				return
			}
			body, err := ExecuteGetOnAdsPage(url)
			if err != nil {
				log.Println("Error while getting response")
				return
			}
			records, err := ParseHttpResp(body, pubName)
			if err != nil {
				log.Println("Error while getting response")
				return
			}
			storage.AddRecordsInDB(records, pubName)
		}(url)
	}
	wg.Wait()
}

//GetPublisherName retrieves the publisher domain from the urls so that later it can be used to store in the db
func GetPublisherName(pubUrl string) (string, error) {
	u, err := url.Parse(pubUrl)
	if err != nil {
		log.Fatal(err)
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
	return body, err
}

func ParseHttpResp(body []byte, pubName string) ([]storage.Record, error) {
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
				fmt.Println(splitOnEachComma)
				fmt.Println("before Comment:", splitBeforeComment)
				fmt.Println("line: ", v)
				fmt.Println("pubName:", pubName)
				continue
			}
			record.Supply_source_domain = domain
			record.Id = strings.TrimSpace(splitOnEachComma[1])
			record.Relationship = relation
			FileData = append(FileData, record)
		}
	}
	fmt.Println(len(FileData))
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
		fmt.Println("default", input)
		return "", errors.New("Invalid relation value")
	}
}
