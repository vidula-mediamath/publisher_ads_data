package main

import("log"
    	"strings"
	"sync"
	"io/ioutil"
	"net/http"
	"net/url"
	"database/sql"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
        )

func main(){
	var wg sync.WaitGroup
	//urls := []string{"https://www.yahoo.com/ads.txt", "https://www.cnn.com/ads.txt", "http://www.nytimes.com/ads.txt", "https://www.nbc.com/ads.txt"}
	urls := []string{"https://www.nbc.com/ads.txt"}
	for _,url := range urls{
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
        	go func(url string) {
                	db,err := storage.GetDbConnection()
			if err != nil{
                                log.Println("DB connection Error.")
                                return
                        }
			// Decrement the counter when the goroutine completes.
                	defer wg.Done()
                	// Fetch the URL.
			pubName, err := GetPublisherName(url)
			if err != nil{
				log.Println("Invalid url")
				return
			}
			body, err := ExecuteGetOnAdsPage(url)
			if err != nil{
				log.Println("Error while getting response")
				return
			}
			ParseHttpResp(body, pubName,  db)
        	}(url)
	}
	wg.Wait()
}


//retrieve the publisher name from the urls so that later it can be used to store in the db
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

func ParseHttpResp(body []byte, pubName string, db *sql.DB) {
var s1[]string = strings.Split(string(body), "\n")
for _,v := range s1{
	var s[]string = strings.Split(v, ",")
	if(len(s) >= 3){
		storage.DbInsert(s, pubName, db)
	}
	// This will only return last error
	}
}
