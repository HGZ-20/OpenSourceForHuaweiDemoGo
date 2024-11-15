// Copyright © 2021 Bin Liu <bin.liu@enmotech.com>

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

func main() {
	connStr := "host=110.41.133.35 port=8000 user=root password=Admin@123456# dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	var date string
	err = db.QueryRow("select current_date ").Scan(&date)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
