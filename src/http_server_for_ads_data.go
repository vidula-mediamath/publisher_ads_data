package main

import ("net/http"
	"log"
	"encoding/json"
	)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	pub_name := r.URL.Path[6:]
	workingURL := map[string]bool {
    		"cnn": true,
    		"nbc": true,
		"nytimes": true,
	}
	if !workingURL[pub_name] {
    		w.Write([]byte("can not produce ads data for this publisher"))
		return
	}
	type Response struct {
    		Supply_source_domain string 
    		Id string	
		Relationship string
	}	
	rows, err := db.Query("SELECT supply_source_domain, id, relationship FROM publisher_ads_data WHERE publisher_name=$1", pub_name)
	if err != nil {
        	log.Fatal(err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
  	if err != nil {
      	log.Fatal(err) 
      	return
  	}
  	count := len(columns)
  	tableData := make([]map[string]interface{}, 0)
  	values := make([]interface{}, count)
  	valuePtrs := make([]interface{}, count)
  	for rows.Next() {
      	for i := 0; i < count; i++ {
          	valuePtrs[i] = &values[i]
      	}
      	rows.Scan(valuePtrs...)
      	entry := make(map[string]interface{})
      	for i, col := range columns {
          	var v interface{}
          	val := values[i]
          	b, ok := val.([]byte)
          	if ok {
              	v = string(b)
          	} else {
              	v = val
          	}
          	entry[col] = v
      	}
      	tableData = append(tableData, entry)
  	}
  	jsonData, err := json.Marshal(tableData)
  	if err != nil {
		log.Fatal(err)
       	return
  	}
	w.Header().Set("Content-Type", "application/json")
        w.Write(jsonData)
}

