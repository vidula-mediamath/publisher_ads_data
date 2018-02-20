package main

import ("testing"
	"net/http"
	"io"
	"fmt"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage")

func TestReadFromAdsFile(t *testing.T){
	storage.GetDbConnection()
	//first start an http server to host files with different content

	//now make a slice of url strings
//	urls := []string{"http://localhost:6080/test", "http://localhost:6080/test/1", "http://localhost:6080/test/2", "https://www.vidulasabnis.com/ads.txt"}
	urls := []string{"http://www.cnn.com/ads.txt"}

	//now call readFromAdsFile
	cTest := make(chan string)
	for _,v := range urls {
		ReadFromAdsFile(v, "test", cTest)
		fmt.Println("after ReadFromAdsFile")
		if "done" != <-cTest {
			t.Error("Test Failed")
		}		
        }
}

//now create an http handler function for the defined http route
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testHandler got called")
	urlParam := r.URL.Path[6:]
	if "1" == urlParam{
		fmt.Println("path must have 1 in it")
        	w.Write(nil)
        }else if "2" == urlParam{
		fmt.Println("path must have 2 in it")
		io.WriteString(w, "hello, world!\n")	
	}else {
		w.Write([]byte("This is an example server.\n"))
	} 
}
