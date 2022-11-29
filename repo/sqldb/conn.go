package sqldb

import "database/sql"

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "username:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	return db
}
