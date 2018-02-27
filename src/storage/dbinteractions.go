package storage

import (
	"database/sql"
	"fmt"
	"log"
)

const (
  host     = "localhost"
  port     = 5432
  user	   = "vsabnis"
  dbname   = "vsabnis"
)

type Record struct{
	Supply_source_domain string
	Id                   string
	Relationship			string
	Created_on			string
	Updated_on			string
}

type DataAccessLayer interface {
  dbQuery(string) []Record
  dbInsert([]Record) error
}

func NewDBConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",host, port, user, dbname)
	db, err := GetDbConnection(psqlInfo)
	if err != nil {
		log.Println("DB connection Error.")
		return nil, err
	}
	return db, err
}

func GetFromDB(pubName string) ([]Record, error) {
	db, err := NewDBConnection()
	defer db.Close()
	if err != nil {
		log.Println("DB connection Error.")
		return nil, err
	}
	tableData, err := dbQuery(pubName, db)
	if err != nil {
		return nil, err
	}
	return tableData, err
}

func AddRecordsInDB(records []Record, pubName string) error {
	db, err := NewDBConnection()
	defer db.Close()
	if err != nil {
		log.Println("DB connection Error.")
		return err
	}
	for _, v := range records {
		dbInsert(v, pubName, db)
		//log.Println(err)
		//todo propagate this error above
	}
	return err
}

//DbQuery will retrieve records from database table for this particular query
func dbQuery(pubName string, db *sql.DB) ([]Record, error) {
	rows, err := db.Query("SELECT supply_source_domain, id, relationship, created_on, updated_on FROM publisher_ads_data WHERE publisher_name=$1", pubName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	tableData := []Record{}
	for rows.Next() {
		entry := Record{}
		if err := rows.Scan(&entry.Supply_source_domain, &entry.Id, &entry.Relationship, &entry.Created_on, &entry.Updated_on); err != nil {
			log.Fatal(err)
		}
		tableData = append(tableData, entry)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tableData, err
}

//DbInsert will insert record in the publisher_ads_data table
func dbInsert(record Record, pubName string, db *sql.DB) error {
	sqlStatement := `insert into publisher_ads_data (publisher_name, supply_source_domain, id, relationship, comment)
values($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, pubName, record.Supply_source_domain, record.Id, record.Relationship, "")
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"my_table_pkey\"" {
			log.Println("")
		}
	}
	return err
}
