package main

import ("testing"
	"io/ioutil"
	"net/http"
	"strings")

var testtable = []struct {
	in  string
	out bool
}{
	{"http://localhost:8080/view/www.cnn.com", true},
	{"http://localhost:8080/view/www.nbc.com", true},
	{"http://localhost:8080/view/www.yahoo.com", true},
	{"http://localhost:8080/view/www.nytimes.com", true},
	{"http://localhost:8080/view/www.abc.com", false},
	{"http://localhost:8080/view/cnn", false},
	{"http://localhost:8080/view/abc", false},
	{"http://localhost:8080/view/", false},
}

func TestHttpRes(t *testing.T) {
	for _, tt := range testtable {
	resp,err := http.Get(tt.in)
	if err != nil {
    		t.Error("test failed for %s", tt.in)
		tt.out = false
		continue
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    		t.Error("test failed while reading response for %s", tt.in)
		tt.out = false
                continue
	}
	if !strings.Contains(string(body), "supply_source_domain"){
		tt.out = false
                continue
	}
	tt.out = true
	}
}
