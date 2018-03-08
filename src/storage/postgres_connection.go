package storage

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

//GetDbConnection will obtain postgres connection and pings the db
func GetDbConnection(connectionString string) (*sql.DB, error) {	
	var db *sql.DB
	db, err := sql.Open("postgres", connectionString)
	if err != nil{		
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, err
}