// +build ignore

package main

import (
	"database/sql"
	"fmt"
)

// Up is executed when this migration is applied
func Up_20141012220956(txn *sql.Tx) {

	fmt.Println("Hello from migration Up_20141012220956 Up!")

	txn.Exec("CREATE TABLE \"user\" (id BIGSERIAL PRIMARY KEY, email TEXT, password TEXT, name TEXT);")

}

// Down is executed when this migration is rolled back
func Down_20141012220956(txn *sql.Tx) {
	 fmt.Println("Hello from migration Down_20141012220956 Down!")

	 txn.Exec("DROP TABLE \"user\"")
}
