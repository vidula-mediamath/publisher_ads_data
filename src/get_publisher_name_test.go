package main

import "testing"

func TestPublisherNameFunc(t *testing.T){
	tables := []struct {
		url string	
		pub_name string 
	}{
		{"https://www.cnn.com/ads.txt", "cnn"},
		{"https://cnn.com/ads.txt", "cnn"},
		{"cnn.com/ads.txt", "cnn"},
		{"cnn", "not a valid url"},
	}

	for _, table := range tables {
		output := GetPublisherName(table.url)
		if output != table.pub_name {
			t.Errorf("Function did not return expected publisher name for url %s", table.url)
		}
	}	
	
}

