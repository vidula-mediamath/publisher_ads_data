package main

import ("testing"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage")

func fakeDbQueryFunc(pubName string) ([]storage.Record, error) {
    return []storage.Record{storage.Record{"www.cnn.com", "21908", "DIRECT", "", ""}, storage.Record{"www.nbc.com", "9010", "DIRECT", "", ""}}, nil
}

func TestGetResponseFromDB(t *testing.T) {
	urlPath := "/view/www.nbc.com"
	
	responseData, err := getResponseFromDB(urlPath, fakeDbQueryFunc)
	if err != nil {
		t.Error(err)
	}
	expectedOutput, _ := fakeDbQueryFunc("abc")
	
	if compareLists(expectedOutput, responseData) != true{
		t.Error("test failed")
	}
}

func fakeDbReturnEmpty(pubName string) ([]storage.Record, error) {
    return []storage.Record{}, nil
}

var testInputs = []struct {
	urlPath string
	expectedFail bool
}{
	{"/check/www.nbc.com", true},
	{"/view/nbc.com", true},
	{"/view/nbc", true},
	{"/view/", true},
}

func TestEmptyResponses(t *testing.T){
	for _,v := range testInputs{
		responseData, err := getResponseFromDB(v.urlPath, fakeDbReturnEmpty)
		if err != nil {
			t.Error(err)
			continue
		}
		expectedOutput := []storage.Record{}
	
		if compareLists(expectedOutput, responseData) != true {
			t.Error("test failed when url path was", v.urlPath)
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