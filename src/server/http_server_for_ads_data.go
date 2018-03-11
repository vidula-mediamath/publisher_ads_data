package main

import ("net/http"
	"strings"
	"errors"
	"encoding/json"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
	)

type apiResponse struct{
	Supply_source_domain	 	string	`json:"supply_source_domain"`
	Id						string	`json:"id"`
	Relationship					string	`json:"relation"`
	Created_on 				string	`json:"created_on"`
	Updated_on 				string	`json:"updated_on"`
}

func main(){
	http.HandleFunc("/view/", viewHandler)
    	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    responseData, err := getResponseFromDB(r.URL.Path, actualDBFunc)
    if err != nil {
	    	w.Write([]byte("can not produce ads data for this publisher"))
        return
    }
    
    httpResponse := []apiResponse{}
	//convert output from db query in to apiResponse format
	for _,v := range responseData{
		respRecord := apiResponse{}
		respRecord.Supply_source_domain = v.Supply_source_domain
		respRecord.Id	=	v.Id
		respRecord.Relationship	= v.Relationship
		respRecord.Created_on = v.Created_on
		respRecord.Updated_on = v.Updated_on
		httpResponse = append(httpResponse, respRecord)
	}

	//convert to json
    jsonData, err := json.Marshal(httpResponse)   
	if err != nil {
        	w.Write([]byte("can not produce ads data for this publisher"))
        return
    }
    // Even though we get different errors, we will show user only one error message.
    if err != nil {
        w.Write([]byte("can not produce ads data for this publisher"))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

type DBReader func(pubName string) ([]storage.Record, error)

func actualDBFunc(pubName string) ([]storage.Record, error) {
	db, err := storage.NewPostgres()
	defer db.Close()
	if err != nil{
		return nil, err
	}
    return db.DBQuery(pubName)
}

func getResponseFromDB(urlPath string, dbQueryFunc DBReader) ([]storage.Record, error) {
    pubName := urlPath[6:]
    tableData, err := dbQueryFunc(pubName)
	return tableData, err
}