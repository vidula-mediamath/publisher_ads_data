package main

import ("testing"
	"fmt"
	)

func TestReadFromAdsFile(t *testing.T) {
	//storage.GetDbConnection()

	urls := []string{"http://localhost:6080/test", "http://localhost:6080/test/1", "http://localhost:6080/test/2", "https://www.vidulasabnis.com/ads.txt"}
//	urls := []string{"http://www.cnn.com/ads.txt"}

	cTest := make(chan string)
	for _,v := range urls {
		go ReadFromAdsFile(v, "test", cTest)
		fmt.Println("after ReadFromAdsFile")
        }
	if "done" != <-cTest {
		t.Error("Test Failed")
	}		
}
