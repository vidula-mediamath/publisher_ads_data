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

//TODO
//write mock db insert or use a test db
	
func TestParseHttpResp(t *testing.T) {
	for i, tt := range testinputs {
        err := ParseHttpResp(tt.in, "test") 
	if err != nil{
		if tt.out == false{
			continue
		}else{
			t.Error("Error in test case number ", i+1)
	}
	}
}
}
