package main

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
)

func TestExecuteGetOnAdsPage(t *testing.T){
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain; charset=CP932")		
		fmt.Fprintln(w, "Hello, \x90\xA2\x8A\x45")
	}))
	defer ts.Close()
	
	body, err := ExecuteGetOnAdsPage(ts.URL)
	if err != nil {
		t.Error(err)
	}
	
	if string(body) != "Hello, 世界"{
		t.Error("Test failed")		
	}
}