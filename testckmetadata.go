package main

import (
	"fmt"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"net/url"
)

func InitClickHouse() (*gorm.DB, error) {
	password := "1qaz2wsx#"
	encodedPassword := url.QueryEscape(password)

	dsn := fmt.Sprintf("clickhouse://default:%s@192.168.0.23:29000/analytics?dial_timeout=10s&read_timeout=20s", encodedPassword)
	return gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
}

func main() {
	db, err := InitClickHouse()
	if err != nil {
		fmt.Println("failed to initialize database, got error", err)
	} else {
		fmt.Println("database initialized successfully", db)
	}

	GetDatabase(db)

	sql := "select * from  unified_log_maas;"

	executeSql(sql, db)

}

func GetDatabase(db *gorm.DB) {
	var databases []string

	db.Raw("show databases").Pluck("name", &databases)

	fmt.Println("databases: =", databases)
}

func executeSql(sql string, db *gorm.DB) {
	rows, err := db.Raw(sql).Rows()

	if err != nil {
		fmt.Println("failed to execute sql, got error", err)
		return
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("columns: =", columns)

}
