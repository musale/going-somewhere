package main

import (
	"net/http"

	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/musale/going-somewhere/golang-postgres/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Info("Error loading .env file")
	}
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	log.WithFields(log.Fields{"inserting fields": true}).Info("Inserting fields")
	var lastInsertID int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertID)
	checkErr(err)
	log.WithFields(log.Fields{"lastInsertID": lastInsertID}).Info("Last insert id")

	log.WithFields(log.Fields{"action": "update"}).Info("Updating userinfo where uid=lastInserID")
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("musale", lastInsertID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	log.WithFields(log.Fields{"action": "change", "affect": affect}).Info("Rows changed")

	http.HandleFunc("/", lib.MessagesPage)
	http.ListenAndServe(":5505", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
