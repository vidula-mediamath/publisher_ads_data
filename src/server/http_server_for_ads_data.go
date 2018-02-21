package main

import ("net/http"
	"log"
	"strings"
	"encoding/json"
	"github.com/vidula-mediamath/publisher_ads_data/src/storage"
	)

func main(){
	http.HandleFunc("/view/", viewHandler)
    	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path, "view"){
		w.Write([]byte("can not produce ads data for this publisher"))
                return
	}
	pubName := r.URL.Path[6:]
	tableData, err := storage.GetFromDB(pubName)
	if err != nil {
		w.Write([]byte("can not produce ads data for this publisher"))
		return
	}

	//convert output from db query in to json
	jsonData, err := json.Marshal(tableData)
        if err != nil {
                log.Fatal(err)
		w.Write([]byte("can not produce ads data for this publisher"))
        }

	w.Header().Set("Content-Type", "application/json")
        w.Write(jsonData)
}
