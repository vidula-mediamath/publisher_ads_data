package main

import("fmt"
        "strings"
	"io/ioutil"
	"net/http"
        )

func main(){
	get_db_connection()
	c := make(chan string)
	urls := []string{"https://www.cnn.com/ads.txt", "http://www.nytimes.com/ads.txt", "https://www.nbc.com/ads.txt"}
	for _,v := range urls{
		go readFromAdsFile(v, c)
	}
	for i:=0; i<len(urls); {
		if "done" == <-c{
			i++
		}
	}
	http.HandleFunc("/view/", viewHandler)
    	http.ListenAndServe(":8080", nil)
}

func GetPublisherName(pubUrl string) (string) {
	var pub_name string
	if strings.Contains(pubUrl, "www") {
		split_pub_name := strings.Split(pubUrl, "www.")
 		pub_name = strings.Split(split_pub_name[1], ".com")[0]	
	} else if strings.Contains(pubUrl, ".com") {
		split_name := strings.Split(pubUrl, ".com")
		if strings.Contains(split_name[0], "://") {
			pub_name = strings.Split(split_name[0], "://")[1] 
		}else{
			pub_name = split_name[0]
		}
	} else {
		return "not a valid url"
		} 		
	return pub_name	
}

func readFromAdsFile(pubUrl string, c chan string) {
	pub_name := GetPublisherName(pubUrl)
	if pub_name == "not a valid url"{
		return
	}
	resp, err := http.Get(pubUrl)
	if err != nil {
			fmt.Println("http called failed")
			return
			}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var s1[]string = strings.Split(string(body), "\n")
	for _,v := range s1{
		var s[]string = strings.Split(v, ",")
		if(len(s) >= 3){
			sendtodb(s, pub_name)
		}
	}
	c <- "done"
}

func sendtodb(array []string, pub_name string) {
	var fixed_length [6]string
	copy(fixed_length[:], array)

	sqlStatement := `insert into publisher_ads_data (publisher_name, supply_source_domain, id, relationship, comment, type1, type2)
values($1, $2, $3, $4, $5, $6, $7)`
	mutex.Lock()
	_, err = db.Exec(sqlStatement, pub_name, fixed_length[0], fixed_length[1], fixed_length[2], fixed_length[3], fixed_length[4], fixed_length[5])
	if err != nil {
		panic(err)
	}
	mutex.Unlock()
	
}

