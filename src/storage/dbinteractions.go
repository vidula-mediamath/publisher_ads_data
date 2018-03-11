package storage

import (
	"database/sql"
	"fmt"
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

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (*Postgres, error) { 
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",host, port, user, dbname)
	db, err := GetDbConnection(psqlInfo)
	if err != nil {
		return nil, err
	}
	
	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) DBQuery(pubName string) ([]Record , error) {
	rows, err := p.db.Query("SELECT supply_source_domain, id, relationship, created_on, updated_on FROM publisher_ads_data WHERE publisher_name=$1", pubName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tableData := []Record{}
	
	for rows.Next() {
		entry := Record{}
		if err := rows.Scan(&entry.Supply_source_domain, &entry.Id, &entry.Relationship, &entry.Created_on, &entry.Updated_on); err != nil {
			return nil, err
		}
		tableData = append(tableData, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tableData, err
}

func (p *Postgres) DBInsert(records []Record, pubName string) []error {
	sqlStatement := `insert into publisher_ads_data (publisher_name, supply_source_domain, id, relationship, comment)
			values($1, $2, $3, $4, $5)`
	var allErrors []error
	for _, record := range records {
		_, err := p.db.Exec(sqlStatement, pubName, record.Supply_source_domain, record.Id, record.Relationship, "")
		if err != nil {
			//check errorcode is for unique key constraint
			if err.Error() == "pq: duplicate key value violates unique constraint \"my_table_pkey\"" {
				continue
			}
		}
		allErrors =append(allErrors, err)	
	}	
	return allErrors
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
