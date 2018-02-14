package main

import ("testing"
	"io/ioutil"
	"net/http"
	"fmt")

func TestHttpRes(t *testing.T){
	resp,err := http.Get("http://localhost:8080/view/cnn")
	if err != nil {
    		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    		t.Error(err)
	}
	fmt.Println(string(body))
}
