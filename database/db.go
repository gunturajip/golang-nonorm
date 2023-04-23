package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	connection = os.Getenv("DB_CONNECTION")
// 	host       = os.Getenv("DB_HOST")
// 	user       = os.Getenv("DB_USERNAME")
// 	password   = os.Getenv("DB_PASSWORD")
// 	port       = os.Getenv("DB_PORT")
// 	dbname     = os.Getenv("DB_DATABASE")
// 	db         *sql.DB
// 	err        error
// )

var (
	db  *sql.DB
	err error
)

func StartDB() {
	// config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	// config := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
	// db, err = sql.Open(connection, config)

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/jobhun_tes")
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	log.Println("connected to database successfully")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	err = db.Close()
	if err != nil {
		log.Fatal("error closing to database :", err)
	}
}
