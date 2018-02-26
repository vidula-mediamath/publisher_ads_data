package main

import ("testing"
		"fmt"
		"github.com/vidula-mediamath/publisher_ads_data/src/storage")

var testinputs = []struct {
        in  []byte
        out []storage.Record
}{
		//test for correct input, happy path	
        {[]byte("www.cnn.com, abc, DIRECT, abc, #abc"), [struct{www.cnn.com abc DIRECT}]},
        //happy path with ReSeller as input instead of RESELLER
        {[]byte("www.cnn.com, pub-2892, ReSeller, abc, #abc"), [{www.cnn.com pub-2892 RESELLER}]},
        //happy path, direct - all lowercase
        {[]byte("www.cnn.com, 21908, direct, abc, #abc"), [{www.cnn.com 21908 DIRECT}]},
        //test for multiple lines present in ads.txt file
        {[]byte("www.cnn.com, 21908, direct, abc, #abc \n www.nbc.com, 9010, Direct, #jkdhsk"), [{www.cnn.com 21908 DIRECT} {www.nbc.com 9010 DIRECT}]},
        //test for invalid supply source domain input
        {[]byte("test, abc, RESELLER"), []},
        //test for no spaces between comma separated values
        {[]byte(".com,abc,RESELLER,abc"), [{.com abc RESELLER}]},
        //test for invalid number and type of inputs
		{[]byte("1234,abc"), []},
		//test for 
		{[]byte(".com,  another, #abc"), []},
		{[]byte("#"), []},
        {[]byte("www.cnn.com,abc,direct,abc,abc,abc,abc,abc"), []},
        {[]byte("www.cnn.com abc DIRECT abc"), [{www.cnn.com abc DIRECT}]},
}
	
func TestParseHttpResp(t *testing.T) {
	for i, tt := range testinputs {
        records, err := ParseHttpResp(tt.in)
        if err != nil{
			fmt.Println(err)
			if tt.out == false{
				continue
			}else{
				t.Error("Error in test case number ", i+1)
			}
		}
        if !records == tt.out {
	        	t.Error("Parsing has error, input was", tt.in)
        }
	}
}
