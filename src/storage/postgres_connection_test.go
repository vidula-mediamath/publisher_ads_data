package storage

import	("testing"
		"fmt")

const (
  testhost     = "localhost"
  testport     = 5432
  testuser	   = "vsabnis"
  testdbname   = "vsabnis"
)


func TestDBConnection(t *testing.T){
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",testhost, testport, testuser, testdbname)
	db,err := GetDbConnection(connectionString)
	if err != nil {
		t.Fatal("Test could not get a db connection")
	}
	_, errCreate := db.Exec("create database test")
	if errCreate != nil {
		t.Error("test failed")
	}
	_, errDrop := db.Exec("drop database test")
	if errDrop != nil {
                t.Error("test failed")
        }
	defer db.Close()
}


