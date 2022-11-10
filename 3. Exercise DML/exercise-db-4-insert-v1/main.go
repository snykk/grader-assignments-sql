package main

import (
	"database/sql"
	"encoding/base64"
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

func Connect(creds *Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

//go:embed insert.sql
var insertStr string

func InsertSQL(dbConn *sql.DB) error {
	_, err := dbConn.Exec(insertStr)
	if err != nil {
		return err
	}

	fmt.Println("success insert data")
	return nil
}

var (
	sqlScript1 = "Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgc3R1ZGVudHMgKAoJCWlkIElOVCwgCgkJZmlyc3RfbmFtZSBWQVJDSEFSKDEwMCksCgkJbGFzdF9uYW1lIFZBUkNIQVIoMTAwKSwKCQlkYXRlX29mX2JpcnRoIERBVEUsCgkJYWRkcmVzcyBWQVJDSEFSKDI1NSksCgkJY2xhc3MgVkFSQ0hBUigxMDApLAoJCXN0YXR1cyBWQVJDSEFSKDEwMCkKCSk="
)

func CreateTable(dbConn *sql.DB) error {
	sqlScript1, _ := base64.StdEncoding.DecodeString(sqlScript1)
	_, err := dbConn.Exec(string(sqlScript1))

	if err != nil {
		return err
	}

	fmt.Println("success create table")
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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS students CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = CreateTable(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = InsertSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
