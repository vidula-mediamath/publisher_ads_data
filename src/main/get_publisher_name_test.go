package main

import "testing"

var tables = []struct {
	url string	//input	
	expectedOutput string	//expected output 
}{
	{"https://www.cnn.com/ads.txt", "www.cnn.com"},
	{"https://cnn.com/ads.txt", "cnn.com"},
	{"cnn.com/ads.txt", ""},
	{"cnn", ""},
}

func TestPublisherNameFunc(t *testing.T){
	for _, table := range tables {
		output,err := GetPublisherName(table.url)
		if err != nil {
			if table.expectedOutput == "" {
				continue
			}else { 
				t.Error("test case failed for input ", table.url)
				continue
		}
		if output != table.expectedOutput {
			t.Errorf("Function did not return expected publisher name for url %s got %s", table.url, output)
		}
	}	
	
}
}

