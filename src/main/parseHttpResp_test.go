package main

import ("testing"
		//"fmt"
		"github.com/vidula-mediamath/publisher_ads_data/src/storage")

var testinputs = []struct {
        in  []byte
        out []storage.Record
}{
		//test for correct input, happy path	
        {[]byte("cnn.com, abc, DIRECT, abc, #abc"), []storage.Record{storage.Record{"cnn.com", "abc", "DIRECT", "", ""}}},
        //happy path with ReSeller as input instead of RESELLER
        {[]byte("cnn.com, pub-2892, ReSeller, abc, #abc"), []storage.Record{storage.Record{"cnn.com", "pub-2892", "RESELLER", "", ""}}},
        //happy path, direct - all lowercase
        {[]byte("cnn.com, 21908, direct, abc, #abc"), []storage.Record{storage.Record{"cnn.com", "21908", "DIRECT", "", ""}}},
        //test for multiple lines present in ads.txt file
        {[]byte("cnn.com, 21908, direct, abc, #abc \nnbc.com, 9010, Direct, #jkdhsk"), 
	        	[]storage.Record{storage.Record{"cnn.com", "21908", "DIRECT", "", ""}, storage.Record{"nbc.com", "9010", "DIRECT", "", ""}}},
        //test for invalid supply source domain input
        {[]byte("test, abc, RESELLER"), []storage.Record{}},
        //test for no spaces between comma separated values
        {[]byte(".com,abc,RESELLER,abc"), []storage.Record{}},
        //test for invalid number and type of inputs
		{[]byte("1234,abc"), []storage.Record{}},
		//test for invalid value of relationship
		{[]byte(".com,  another, #abc"), []storage.Record{}},
		//test for line with #
		{[]byte("#"), []storage.Record{}},
		//test for line with valid data containing more values after comment
        {[]byte("cnn.com,abc,direct,abc,abc,abc,abc,abc"), []storage.Record{storage.Record{"cnn.com", "abc", "DIRECT", "", ""}}},
        //test for input that is not comma separated
        {[]byte("cnn.com abc DIRECT abc"), []storage.Record{}},
        //test for empty file
        {[]byte(""), []storage.Record{}},
        //test for different encodings
        {[]byte("cnn.com, こんにちは,RESELLER,abc"), []storage.Record{storage.Record{"cnn.com", "こんにちは", "RESELLER", "", ""}}},
        //test for wide range of characters
        {[]byte("cnn.com, #@$@&*^€¢£¥^€, RESELLER, abc"), []storage.Record{}},
        //test for supply domain url with www.
        {[]byte("www.cnn.com,abc,direct,abc,abc,abc,abc,abc"), []storage.Record{storage.Record{"www.cnn.com", "abc", "DIRECT", "", ""}}},
}
	
func TestParseHttpResp(t *testing.T) {
	for i, tt := range testinputs {
        records,_ := parseHttpResp(tt.in)
        //fmt.Println(records)
        equal := compareLists(records, tt.out)
        if !equal {
	        	t.Error("Parsing has error, test number was", i+1)
        }
	}
}

func compareLists(slice1 []storage.Record, slice2 []storage.Record) bool {
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
