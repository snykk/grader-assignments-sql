package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "embed"

	_ "github.com/lib/pq"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type TableCheck struct {
	ColumnName      string `sql:"column_name"`
	OrdinalPosition int    `sql:"ordinal_position"`
	IsNullable      string `sql:"is_nullable"`
	DataType        string `sql:"data_type"`
	CharLength      int    `sql:"character_maximum_length"`
}

func Connect(creds *Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

//go:embed create-table.sql
var command string

func CreateTable(db *sql.DB) error {
	res, err := db.Exec(command)
	if err != nil {
		return err
	}

	fmt.Println(res)
	fmt.Println("Table created successfully")

	return nil
}

func main() {
	dbCredential := Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateTable(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
