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

type Report struct {
	ID       int    `sql:"id"`
	Fullname string `sql:"fullname"`
	Class    string `sql:"class"`
	Status   string `sql:"status"`
	Study    string `sql:"study"`
	Score    int    `sql:"score"`
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

//go:embed select.sql
var queryStr string

func QueryJoinSQL(db *sql.DB) ([]Report, error) {
	var res []Report

	rows, err := db.Query(queryStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r Report
		err := rows.Scan(&r.ID, &r.Fullname, &r.Class, &r.Status, &r.Study, &r.Score)
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}

	return res, nil
}

var (
	CreateSQL = `Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgc3R1ZGVudHMgKAoJCWlkIFNFUklBTCBQUklNQVJZIEtFWSwgZnVsbG5hbWUgVkFSQ0hBUigyNTUpIE5PVCBOVUxMLCBkYXRlX29mX2JpcnRoIERBVEUgTk9UIE5VTEwsCgkJY2xhc3MgVkFSQ0hBUigyNTUpIE5PVCBOVUxMLCBzdGF0dXMgVkFSQ0hBUig1MCkgTk9UIE5VTEwKCSk7CglDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyByZXBvcnRzICgKCQlpZCBTRVJJQUwgUFJJTUFSWSBLRVksIHN0dWRlbnRfaWQgSU5UIE5PVCBOVUxMLCBzdHVkeSBWQVJDSEFSKDI1NSkgTk9UIE5VTEwsCgkJc2NvcmUgSU5UIE5PVCBOVUxMLCBGT1JFSUdOIEtFWSAoc3R1ZGVudF9pZCkgUkVGRVJFTkNFUyBzdHVkZW50cyhpZCkKCSk7`
	InsertSQL = `SU5TRVJUIElOVE8gc3R1ZGVudHMgKGZ1bGxuYW1lLCBkYXRlX29mX2JpcnRoLCBjbGFzcywgc3RhdHVzKSAKCVZBTFVFUyAoJ0pvaG4gRG9lJywgJzIwMDAtMDEtMDEnLCAnMUEnLCAnYWN0aXZlJyksCgkoJ0phbmUgV2lsbHknLCAnMjAwMy0wMS0wMScsICcxQScsICdhY3RpdmUnKSwKCSgnSm9obiBTbWl0aCcsICcyMDAyLTAzLTAxJywgJzFBJywgJ2luYWN0aXZlJyksCgkoJ0JvYiBBYnJhbXMnLCAnMjAwMy0wNC0wMScsICcxQicsICdhY3RpdmUnKTsKSU5TRVJUIElOVE8gcmVwb3J0cyAoc3R1ZGVudF9pZCwgc3R1ZHksIHNjb3JlKQpWQUxVRVMgKDEsICdNYXRoJywgOTApLCAoMSwgJ0VuZ2xpc2gnLCA4MCksICgxLCAnU2NpZW5jZScsIDcwKSwgKDEsICdJbmRvbmVzaWEnLCA3MCksCgkoMiwgJ01hdGgnLCA1NSksICgyLCAnRW5nbGlzaCcsIDgwKSwgKDIsICdTY2llbmNlJywgNjEpLCAoMiwgJ0luZG9uZXNpYScsIDcwKSwKCSgzLCAnTWF0aCcsIDkwKSwgKDMsICdFbmdsaXNoJywgODApLCAoMywgJ1NjaWVuY2UnLCA3MCksICgzLCAnSW5kb25lc2lhJywgNzApLAoJKDQsICdNYXRoJywgNjUpLCAoNCwgJ0VuZ2xpc2gnLCAzMCksICg0LCAnU2NpZW5jZScsIDQwKSwgKDQsICdJbmRvbmVzaWEnLCA1MCk7`
)

func SQLExecute(db *sql.DB, insertSQL string) error {
	CreateSQL, _ := base64.StdEncoding.DecodeString(CreateSQL)
	_, err := db.Exec(string(CreateSQL))
	if err != nil {
		fmt.Println("error create table")
		return err
	}

	fmt.Println("create table success")

	insert, _ := base64.StdEncoding.DecodeString(insertSQL)
	_, err = db.Exec(string(insert))
	if err != nil {
		fmt.Println("error insert data")
		return err
	}

	fmt.Println("insert data success")

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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS students, reports CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn, InsertSQL)
	if err != nil {
		log.Fatal(err)
	}

	res, err := QueryJoinSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range res {
		fmt.Println(s)
	}
}
