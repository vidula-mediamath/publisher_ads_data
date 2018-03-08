package main

import "testing"

var testtable = []struct {
        in  string
        out bool
}{
        {"https://www.yahoo.com/ads.txt", true},
        {"https://yahoo.com/ads.txt", true},
        {"yahoo.com/ads.txt", false},
        {"vidulasabnis.com/ads.txt", false},
	{"https://golang.org/ads.txt", false},
}

func TestExecuteGetOnAdsPage(t *testing.T){
	for _, tt := range testtable {
		_,err := executeGetOnAdsPage(tt.in)
		if err != nil {
			if tt.out == false{
				continue	
			}else{
			t.Error("Test failed for ", tt.in)
			continue
		}
	}
}
}

