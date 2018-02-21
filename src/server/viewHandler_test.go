package main

import ("testing"
	"net/http/httptest"
	"net/http")

func TestViewHandler(t *testing.T){
	inputs := []string{"http://localhost:8080/view/www.yahoo.com",
		"http://localhost:8080/view/www.vidulasabnis.com",
		"http://localhost:8080/view/",
		"http://localhost:8080/"}
	for _, v := range inputs{
	req, _ := http.NewRequest("GET", v, nil)
    	w := httptest.NewRecorder()
    	viewHandler(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("page didn't return %v", http.StatusOK)
    }
}	
}

