package storage

import	"testing"

func TestDBConnection(t *testing.T){
	get_db_connection()
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


