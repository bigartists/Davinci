package main

import (
	"davinci/pkg/SqlUtils"
	"fmt"
)

func main() {
	dbType := "mysql"
	user := "root"
	password := "root"
	address := "0.0.0.0"
	port := "3306"

	db, err := SqlUtils.GetDatabaseInstance(dbType, user, password, address, port)
	if err != nil {
		fmt.Println("Error creating database instance:", err)
		return
	}

	databases, err := db.GetSourceDatabases()
	if err != nil {
		fmt.Println("Error getting databases:", err)
		return
	}

	fmt.Println("Databases:", databases)

	tables, err := db.GetSourceTables("davinci3")
	if err != nil {
		fmt.Println("Error getting tables:", err)
		return
	}
	fmt.Println("Tables:", tables)
	//

	columns, err := db.GetSourceTableColumns("davinci3", "source")
	if err != nil {
		fmt.Println("Error getting columns:", err)
		return
	}
	fmt.Println("Columns:", columns)
	//

	results, err := db.ExecuteSql("SELECT * FROM analysis.channel")
	if err != nil {
		fmt.Println("Error executing SQL:", err)
		return
	}

	fmt.Println("Results:", results)
}
