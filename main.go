package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id       int
	name     string
	username string
}

func main() {
	sqlMigration()
	insertData()
	readData()
}

func insertData() {
	var db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO user (name, username) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	statement.Exec("Dante Ramos", "dante123")
}

func readData() {
	var db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, username FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user = user{}

	fmt.Printf("|%-6s|%-15s|%-15s|\n", "id", "Name", "User name")
	fmt.Println("________________________________________")
	for rows.Next() {
		rows.Scan(&user.id, &user.name, &user.username)
		fmt.Printf("|%-6d|%-15s|%-15s|\n", user.id, user.name, user.username)
	}
}

func sqlMigration() {
	var file, err = os.Open("./migration.sql")
	if err != nil {
		log.Fatal("File not found")
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatal(err)
	}
}
