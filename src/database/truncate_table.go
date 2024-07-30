package database

import (
	"fmt"
)

func TruncateTable(table string) {

	sql := fmt.Sprintf("TRUNCATE TABLE %v", table)

	ADPT.Exec(sql)
}
