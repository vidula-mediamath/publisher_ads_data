package main

import ("testing"
	"encoding/json"
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
	expectedJson,_ := json.Marshal(expectedOutput)
	
	if string(expectedJson) != string(responseData){
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
		if err != nil && v.expectedFail != true{
			t.Error(err)
		}else{
			t.Log(err)
			continue
		}
		
		//currently with given inputs, control never reaches this point
		expectedOutput, _ := fakeDbReturnEmpty("abc")
		expectedJson,_ := json.Marshal(expectedOutput)
	
		if string(expectedJson) != string(responseData){
			t.Error("test failed when url path was", v.urlPath)
		}
	}
}