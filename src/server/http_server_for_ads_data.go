package main

import ("net/http"
	"log"
	"encoding/json"
	"database/sql"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
	)

func main(){
	http.HandleFunc("/view/", viewHandler)
    	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pubName := r.URL.Path[6:]
	db, err := storage.GetDbConnection()
	if err != nil{
		w.Write([]byte("can not produce ads data for this publisher"))
		return
		}
	jsonOutput, err := getFromDb(pubName, db)
	if err != nil {
		w.Write([]byte("can not produce ads data for this publisher"))
	}
	w.Header().Set("Content-Type", "application/json")
        w.Write(jsonOutput)

}

func getFromDb(pubName string, db *sql.DB)([]byte,  error){
	tableData,err := storage.DbQuery(pubName, db)
	if err!= nil {
		return nil, err
	}
  	jsonData, err := json.Marshal(tableData)
  	if err != nil {
		log.Fatal(err)
  	}
	return jsonData, err
}
