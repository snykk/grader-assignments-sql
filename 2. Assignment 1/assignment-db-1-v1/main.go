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

type TableCheck struct {
	ColumnName      string `sql:"column_name"`
	OrdinalPosition int    `sql:"ordinal_position"`
	IsNullable      string `sql:"is_nullable"`
	DataType        string `sql:"data_type"`
}

func Connect(creds *Credential) (*sql.DB, error) {
	// this is only an example, please modify it to your need
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

//go:embed add_column_users.sql
var alterAddStr string

//go:embed drop_column_users.sql
var alterDropStr string

func AlterSQL(db *sql.DB) error {
	_, err := db.Exec(alterAddStr)
	if err != nil {
		return err
	}

	fmt.Println("success alter add column")

	_, err = db.Exec(alterDropStr)
	if err != nil {
		return err
	}

	fmt.Println("success alter drop column")

	return nil
}

//go:embed create_table_presences.sql
var createStr string

func CreateSQL(db *sql.DB) error {
	_, err := db.Exec(createStr)
	if err != nil {
		return err
	}

	fmt.Println("success create table")

	return nil
}

//go:embed drop_table_attendances.sql
var dropStr string

func DropSQL(db *sql.DB) error {
	_, err := db.Exec(dropStr)
	if err != nil {
		return err
	}

	fmt.Println("success drop table")

	return nil
}

var (
	sqlScript  = `Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgYXR0ZW5kYW5jZXMgKAoJCWlkIFNFUklBTCBQUklNQVJZIEtFWSwKCQl1c2VyX2lkIElOVCBOT1QgTlVMTCwKCQlzdGF0dXMgSU5UIE5PVCBOVUxMICk=`
	sqlScript2 = `Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgdXNlcnMgKAoJCWlkIElOVEVHRVIgTk9UIE5VTEwsCgkJZnVsbG5hbWUgVkFSQ0hBUigyNTUpIE5PVCBOVUxMLAoJCWVtYWlsIFZBUkNIQVIoMjU1KSBOT1QgTlVMTCwKCQlwYXNzd29yZCBWQVJDSEFSKDI1NSkgTk9UIE5VTEwsCgkJcm9sZSBWQVJDSEFSKDEwMCkKKQ==`
)

func SQLExecute(db *sql.DB) error {
	sqlScript, _ := base64.StdEncoding.DecodeString(sqlScript)
	_, err := db.Exec(string(sqlScript))
	if err != nil {
		return err
	}

	sqlScript2, _ := base64.StdEncoding.DecodeString(sqlScript2)
	_, err = db.Exec(string(sqlScript2))
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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS users, presence, attendance CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = AlterSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = DropSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
