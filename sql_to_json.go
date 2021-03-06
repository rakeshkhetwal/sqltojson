package sqltojson	

import (
	"database/sql"
)

//converting sql rows into json data
func SqlToJson(db *sql.DB, query string) (error,[]map[string]interface {}){
	// prepared statement for repeated use.
	stmt, err := db.Prepare(query)
	if err != nil {
		return err, nil
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err, nil
	}else {
		defer rows.Close()
		columns, err := rows.Columns()
		if err != nil {
			return err, nil
		}
		count := len(columns)
		// creating query data map
		queryData := make([]map[string]interface{}, 0)
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
				//conversion in bytes of values
				b, ok := val.([]byte)
				
				if ok {
					v = string(b)
					
				} else {
					// for nil data
					v = val
				}
				entry[col] = v
			}
			//appending entry into array
			queryData = append(queryData, entry)
		}

		return err, queryData
		
	}
	return err, nil
}
