package main

import (
	"davinci/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//func main() {
//	dsn := config.SysYamlconfig.Database.Dsn
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	databases := GetDatabases(db)
//	for _, database := range databases {
//		fmt.Println("databases: =", database)
//	}
//
//	tables := GetTables(db)
//
//	for _, table := range tables {
//		fmt.Println("tables: =", table)
//	}
//
//}
//
//func GetDatabases(db *gorm.DB) []string {
//	var databases []string
//
//	db.Raw("show databases").Pluck("davinci3", &databases)
//
//	return databases
//}
//
//func GetTables(db *gorm.DB) []string {
//	var tables []string
//
//	db.Raw("show tables").Pluck("davinci3", &tables)
//
//	return tables
//}

type RowData map[string]interface{}
type RowsData []RowData

func main() {
	dsn := config.SysYamlconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbName := "analysis"

	err = db.Exec("USE " + dbName).Error

	if err != nil {
		log.Fatal(err)
	}

	sql := "SELECT * FROM analysis.channel"

	rows, err := db.Raw(sql).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("columns:=", columns)

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	rowData := make(RowData)
	rowsData := make(RowsData, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("Error canning row:", err)
			return
		}

		for i, col := range values {
			if b, ok := col.([]byte); ok {
				rowData[columns[i]] = string(b)
			} else {
				rowData[columns[i]] = col
			}
			//fmt.Print("key:=", columns[i], " value:=", col, " ")
			//fmt.Printf("key:=%s value:=%v type:%s ", columns[i], rowData[columns[i]], reflect.TypeOf(rowData[columns[i]]))
			//fmt.Printf(" type:%s ", reflect.TypeOf(rowData[columns[i]]))
		}

		rowsData = append(rowsData, rowData)

	}

	// 检查是否有错误发生在迭代过程中
	if err = rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
	}

	//var result []map[string]interface{}
	//err = db.Raw(sql).Scan(&result).Error
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//
	//
	//for _, row := range result {
	//	for key, value := range row {
	//		log.Println(key, value)
	//	}
	//	log.Println(".........................")
	//}

	//databases := GetDatabases(db)
	//for _, database := range databases {
	//	fmt.Println("databases: =", database)
	//}
	//
	//tables := GetTables(db)
	//
	//for _, table := range tables {
	//	fmt.Println("tables: =", table)
	//}

}

//func GetDatabases(db *gorm.DB) []string {
//	var databases []string
//
//	db.Raw("show databases").Pluck("davinci3", &databases)
//
//	return databases
//}
//
//func GetTables(db *gorm.DB) []string {
//	var tables []string
//
//	db.Raw("show tables").Pluck("davinci3", &tables)
//
//	return tables
//}
