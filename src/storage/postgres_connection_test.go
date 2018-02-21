package storage

import	"testing"

func TestDBConnection(t *testing.T){
	db,err := GetDbConnection()
	if err != nil {
		t.Fatal("Test could not get a db connection")
	
	_, err := db.Exec("create database test")
	if err != nil {
		t.Error("test failed")
	}
	_, err1 := db.Exec("drop database test")
	if err1 != nil {
                t.Error("test failed")
        }
	defer db.Close()
}


