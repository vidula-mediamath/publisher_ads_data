package storage

import (
  "sync"
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

var mutex sync.Mutex
var db *sql.DB
var err error

const (
  host     = "localhost"
  port     = 5432
  user	   = "vsabnis"
  dbname   = "vsabnis"
)

func GetDbConnection() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",host, port, user, dbname)

db, err = sql.Open("postgres", psqlInfo)
if err != nil {
  panic(err)
}
err = db.Ping()
if err != nil {
  panic(err)
}
fmt.Println("Successfully connected!")
}
