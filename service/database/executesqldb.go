package database

import (
	"fmt"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) ExecuteSQLDB(code string) () {
	// exec the sql code and use fmt to print the output
	_, err := db.c.Exec(code)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Executed SQL code " + code)

	return
}
