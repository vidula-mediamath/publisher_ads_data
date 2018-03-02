package main

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
)

func TestExecuteGetforDiffEncoding(t *testing.T){
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain; charset=CP932")		
		fmt.Fprintln(w, "Hello, \x90\xA2\x8A\x45")
	}))
	defer ts.Close()
	
	body, err := ExecuteGetOnAdsPage(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(body))
	if string(body) != "Hello, 世界" {
		t.Error("Test failed")		
	}
}

func TestExecuteGetUtf8Encoding(t *testing.T){
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain; charset=utf-8")		
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()
	
	body, err := ExecuteGetOnAdsPage(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println(string(body))
	if string(body) != "Hello, client" {
		t.Error("Test failed")		
	}
}