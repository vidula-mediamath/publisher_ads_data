package main

import "testing"

func TestReadFromUrl(t *testing.T){
	get_db_connection()
	c := make(chan string)	
	//test 1 with url "http://www.tripadvisor.com/ads.txt"
	//test 2 with url "http://tripadvisor.com/ads.txt"
	var urls = []string{"http://www.tripadvisor.com/ads.txt", "http://tripadvisor.com/ads.txt", "tripadvisor/ads.txt"}

	for _, url := range urls {
		readFromAdsFile(url, c)

		if "done" == <-c {

			_, err := db.Query("select publisher_name, supply_source_domain, id, relationship from public.publisher_ads_data where publisher_namex = $1 limit 1", "tripadvisor")  
			if err != nil{
				t.Error("test failed")
			}		

		} else {
			t.Error("scanning url failed")
		}
	}
	db.Close()
}

