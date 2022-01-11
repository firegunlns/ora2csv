package main

import (
	"database/sql"
	"fmt"

	go_ora "github.com/sijms/go-ora/v2"
)

// ora2csv
// dump a table from oracle database to csv
// command line usage：
// ora2csv tabble/view
// -h host/ip(default 127.0.0.1) of oracle database server
// -p tcp/ip port(default 1521) of oracle database server
// -q sql query
func ora2csv() {
	fmt.Printf("ora2csv\n")
	server := "8.142.171.235"
	port := 1521
	inst := "orcl"
	table := "BMSQL_CONFIG"
	user := "openroad"
	password := "Q1w2e3r4"

	databaseURL := go_ora.BuildUrl(server, port, inst, user, password, nil)
	conn, err := sql.Open("oracle", databaseURL)
	// check for err
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sql_str := "select count(1) from all_tab_columns where Table_Name = 'BMSQL_CONFIG' "
	rows, err := conn.Query(sql_str)
	// check for err
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var test test1
	rows.Next()

	var cnt int
	rows.Scan(&cnt)

	row := make([]string, cnt)

	rows1, err1 := conn.Query("select * from " + table)
	if err1 != nil {
		fmt.Println(err.Error())
		return
	}

	for rows1.Next() {
		rows1.Scan(&row)
		fmt.Printf("row is %s\n ", row[0])
	}
}

func mysql2csv() {
	fmt.Printf("mysql2csv\n")
	server := "127.0.0.1"
	port := 1521
	inst := "orcl"
	table := "BMSQL_CONFIG"
	user := "openroad"
	password := "Q1w2e3r4"

	databaseURL := go_ora.BuildUrl(server, port, inst, user, password, nil)
	conn, err := sql.Open("oracle", databaseURL)
	// check for err
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sql_str := "select count(1) from all_tab_columns where Table_Name = 'BMSQL_CONFIG' "
	rows, err := conn.Query(sql_str)
	// check for err
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var test test1
	rows.Next()

	var cnt int
	rows.Scan(&cnt)

	row := make([]string, cnt)

	rows1, err1 := conn.Query("select * from " + table)
	if err1 != nil {
		fmt.Println(err.Error())
		return
	}

	for rows1.Next() {
		rows1.Scan(&row)
		fmt.Printf("row is %s\n ", row[0])
	}
}

func test1() {

}

func main() {
}
