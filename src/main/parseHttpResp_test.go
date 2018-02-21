package main

import "testing"

var testinputs = []struct {
        in  []byte
        out bool
}{
        {[]byte("www.cnn.com, abc, abc, abc"), true},
        {[]byte("test, abc, abc"), true},
        {[]byte("1234,abc,abc,abc"), true},
	{[]byte("1234,abc"), true},
	{[]byte("test,  another"), false},
//	{[]byte(12939), false},
        {[]byte("www.cnn.com,abc,abc,abc,abc,abc,abc,abc"), false},
        {[]byte("www.cnn.com abc abc abc"), false},
}
	
func TestParseHttpResp(t *testing.T) {
	for i, tt := range testinputs {
        err := ParseHttpResp(tt.in, "test") 
	if err != nil{
		t.Error("Error in test case number %d", i+1)
		tt.out = false
	}
	tt.out = true
	}
}
