package main

import ("testing"
        "net/http"
        )

func TestHttpResForPublishers(t *testing.T){
	ch := make(chan string)
	go getResponse("http://cnn.com/ads.txt", ch, t)
	go getResponse("http://www.cnn.com/ads.txt", ch, t)
	go getResponse("cnn.com/ads.txt", ch, t)
	go getResponse("http://nbc.com/ads.txt", ch, t)
	go getResponse("http://nytimes.com/ads.txt", ch, t)
	for i := 0; i < 5; {
		if "done" == <-ch{
		i++
		}
   	t.Log("all outputs received")
	}
}
	 
func getResponse(url string, c chan string, t *testing.T){
	_,err := http.Get(url)
        if err != nil {
                t.Error(err)
        }
	c <- "done"	
}
	
