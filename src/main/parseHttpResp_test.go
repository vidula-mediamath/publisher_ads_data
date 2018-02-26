package main

import ("testing"
		//"fmt"
		"github.com/vidula-mediamath/publisher_ads_data/src/storage")

var testinputs = []struct {
        in  []byte
        out []storage.Record
}{
		//test for correct input, happy path	
        {[]byte("www.cnn.com, abc, DIRECT, abc, #abc"), []storage.Record{storage.Record{"www.cnn.com", "abc", "DIRECT"}}},
        //happy path with ReSeller as input instead of RESELLER
        {[]byte("www.cnn.com, pub-2892, ReSeller, abc, #abc"), []storage.Record{storage.Record{"www.cnn.com", "pub-2892", "RESELLER"}}},
        //happy path, direct - all lowercase
        {[]byte("www.cnn.com, 21908, direct, abc, #abc"), []storage.Record{storage.Record{"www.cnn.com", "21908", "DIRECT"}}},
        //test for multiple lines present in ads.txt file
        {[]byte("www.cnn.com, 21908, direct, abc, #abc \n www.nbc.com, 9010, Direct, #jkdhsk"), 
	        	[]storage.Record{storage.Record{"www.cnn.com", "21908", "DIRECT"}, storage.Record{"www.nbc.com", "9010", "DIRECT"}}},
        //test for invalid supply source domain input
        {[]byte("test, abc, RESELLER"), []storage.Record{}},
        //test for no spaces between comma separated values
        {[]byte(".com,abc,RESELLER,abc"), []storage.Record{storage.Record{".com", "abc", "RESELLER"}}},
        //test for invalid number and type of inputs
		{[]byte("1234,abc"), []storage.Record{}},
		//test for invalid value of relationship
		{[]byte(".com,  another, #abc"), []storage.Record{}},
		//test for line with #
		{[]byte("#"), []storage.Record{}},
		//test for line with valid data containing more values after comment
        {[]byte("www.cnn.com,abc,direct,abc,abc,abc,abc,abc"), []storage.Record{storage.Record{"www.cnn.com", "abc", "DIRECT"}}},
        //test for input that is not comma separated
        {[]byte("www.cnn.com abc DIRECT abc"), []storage.Record{}},
}
	
func TestParseHttpResp(t *testing.T) {
	for _, tt := range testinputs {
        records,_ := ParseHttpResp(tt.in)
        equal := compareList(records, tt.out)
        if !equal {
	        	t.Error("Parsing has error, input was", tt.in)
        }
	}
}

func compareList(slice1 []storage.Record, slice2 []storage.Record) bool {
	if len(slice1) != len(slice2) {
			return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
			}
		}
	return true
}	

