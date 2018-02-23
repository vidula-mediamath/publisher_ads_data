package storage

import (
	"database/sql"
	//	"fmt"
	"log"
)

type Record struct {
	Supply_source_domain string
	Id                   string
	Relationship         string
}

func NewDBConnection() (*sql.DB, error) {
	db, err := GetDbConnection()
	if err != nil {
		log.Println("DB connection Error.")
		return nil, err
	}
	return db, err
}

func GetFromDB(pubName string) ([]map[string]interface{}, error) {
	db, err := GetDbConnection()
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
	db, err := GetDbConnection()
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
func dbQuery(pubName string, db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT * FROM publisher_ads_data WHERE publisher_name=$1", pubName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
		return nil, err
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
	return tableData, err
}

//DbInsert will insert record in the publisher_ads_data table
func dbInsert(record Record, pubName string, db *sql.DB) error {
	//fmt.Println(record)
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
