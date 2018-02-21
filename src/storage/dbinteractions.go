package storage

import ("log"
	"strings"
	"database/sql"
	)

//DbQuery will retrieve records from database table for this particular query
func DbQuery(pubName string, db *sql.DB) ([]map[string]interface{}, error){
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
func DbInsert(input []string, pubName string, db *sql.DB) error {
	for _, v := range input{
		v = strings.TrimSpace(v)
	}
		
	var fixed_length [6]string
        copy(fixed_length[:], input)
        sqlStatement := `insert into publisher_ads_data (publisher_name, supply_source_domain, id, relationship, comment, type1, type2)
values($1, $2, $3, $4, $5, $6, $7)`
        _, err := db.Exec(sqlStatement, pubName, fixed_length[0], fixed_length[1], fixed_length[2], fixed_length[3], fixed_length[4], fixed_length[5])
        if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"my_table_pkey\"" {
			log.Println(err)
		}
        }
	return err
}
