package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

type col_info struct {
	name   string
	typ    string
	isnull bool
	key    string
	defval string
	extra  string
}

func mysql2csv() {
	fmt.Printf("mysql2csv\n")
	server := "127.0.0.1"
	port := 3306
	inst := "openroad"
	table := "commodity"
	user := "root"
	password := "11111111"

	//databaseURL := go_ora.BuildUrl(server, port, inst, user, password, nil)
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, inst))
	// check for err
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sql_str := "desc " + table
	rows, err := conn.Query(sql_str)
	// check for err
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cols := make([]col_info, 0)
	for rows.Next() {
		var info col_info
		rows.Scan(&info.name, &info.typ, &info.isnull, &info.key, &info.defval, &info.extra)
		cols = append(cols, info)
	}

	for i := 0; i < len(cols)-1; i++ {
		fmt.Printf("%s,", cols[i].name)
	}
	fmt.Println(cols[len(cols)-1].name)

	row := make([]string, len(cols))
	row1 := make([]interface{}, len(cols))
	for i := 0; i < len(row); i++ {
		row1[i] = &row[i]
	}

	rows1, err1 := conn.Query("select * from " + table + " limit 10")
	if err1 != nil {
		fmt.Println(err.Error())
		return
	}

	for rows1.Next() {
		rows1.Scan(row1...)
		for i := 0; i < len(row)-1; i++ {
			fmt.Printf("%s,", row[i])
		}
		fmt.Println(row[len(row)-1])
	}
}

func main() {
	mysql2csv()
}
