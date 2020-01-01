package gw_s3_handler

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func TestDB() {
	dbConf := ReadDBConnConfig()
	connFmt := "host=%s port=%s dbname=%s user=%s password=%s sslmode=disable"
	connStr := fmt.Sprintf(connFmt, dbConf.Host, dbConf.Port, dbConf.Name, dbConf.User, dbConf.Password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	rows, err := db.Query("SELECT id, username FROM auth_user ORDER BY id LIMIT 10")
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	var (
		id       int
		username string
	)
	for rows.Next() {
		err := rows.Scan(&id, &username)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%d --- %s\n", id, username)
	}
}
