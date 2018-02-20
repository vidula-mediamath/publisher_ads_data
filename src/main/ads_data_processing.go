package main

import("log"
	"fmt"
        "strings"
	"io/ioutil"
	"net/http"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
        )

//this program contains parsing the ads data file and storing data in postgres
func main(){
	storage.GetDbConnection()
	c := make(chan string)
	urls := []string{"https://www.yahoo.com/ads.txt", "https://www.cnn.com/ads.txt", "http://www.nytimes.com/ads.txt", "https://www.nbc.com/ads.txt"}
	for _,v := range urls{
		pubName := GetPublisherName(v)
		go ReadFromAdsFile(v, pubName, c)
	}
	for i:=0; i<len(urls); {
		if "done" == <-c{
			i++
		}
	}
}

//retrieve the publisher name from the urls so that later it can be used to store in the db
func GetPublisherName(pubUrl string) (string) {
	var pubName string
	if strings.Contains(pubUrl, "www") {
		s1 := strings.Split(pubUrl, "www.")
 		pubName = strings.Split(s1[1], ".com")[0]	
	} else if strings.Contains(pubUrl, ".com") {
		s2 := strings.Split(pubUrl, ".com")
		if strings.Contains(s2[0], "://") {
			pubName = strings.Split(s2[0], "://")[1] 
		}else{
			pubName = s2[0]
		}
	} else {
		return "not a valid url"
		} 		
	return pubName	
}

func ReadFromAdsFile(pubUrl string, pubName string, c chan string) {
	defer func() {
	fmt.Println("inside defer function")
        if err := recover(); err != nil {
            log.Println("parsing a publisher file failed:", err)
        }
	fmt.Println("push to channel")
	c <- "done"
	fmt.Println("after push to channel")
    	}()
	
	resp, err := http.Get(pubUrl)
	fmt.Println("got http GET")
	if err != nil {
			fmt.Println("got http error")
			log.Println(err)
			return
			}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("got response body")
	if err != nil {
                        log.Println(err)
                        return
                        }
	var s1[]string = strings.Split(string(body), "\n")
	for _,v := range s1{
		var s[]string = strings.Split(v, ",")
		if(len(s) >= 3){
			fmt.Println("about to build db records")
			storage.DbInsert(s, pubName)
		}
	}
}
