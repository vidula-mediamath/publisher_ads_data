package main

import ("net/http"
	"log"
	"encoding/json"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
	)

func main(){
	storage.GetDbConnection()
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/test/", testHandler)
    	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pubName := r.URL.Path[6:]
	//jsonOutput, err := getFromDb(pubName, storage.DbQuery)
	jsonOutput, err := getFromDb(pubName)
	if err != nil {
		w.Write([]byte("can not produce ads data for this publisher"))
	}
	w.Header().Set("Content-Type", "application/json")
        w.Write(jsonOutput)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
        w.Write(nil)
	}
	
//type QueryDoer func(pubName string) []map[string]interface{}

//func getFromDb(pubName string, dbQuery QueryDoer) (jsonData []byte, err error){		

func getFromDb(pubName string)([]byte,  error){
	tableData,err := storage.DbQuery(pubName)
	if err!= nil {
		return nil, err
	}
  	jsonData, err := json.Marshal(tableData)
  	if err != nil {
		log.Fatal(err)
  	}
	return jsonData, err
}
