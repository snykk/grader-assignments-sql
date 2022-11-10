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

//go:embed insert_students.sql
var insertStr string

func InsertSQL(db *sql.DB) error {
	_, err := db.Exec(insertStr)
	if err != nil {
		return err
	}

	fmt.Println("success insert data")
	return nil
}

//go:embed update_data_teachers.sql
var updateStr string

func UpdateSQL(db *sql.DB) error {
	_, err := db.Exec(updateStr)
	if err != nil {
		return err
	}

	fmt.Println("success update data")
	return nil
}

//go:embed delete_data_teachers.sql
var deleteStr string

func DeleteSQL(db *sql.DB) error {
	_, err := db.Exec(deleteStr)
	if err != nil {
		return err
	}

	fmt.Println("success delete data")
	return nil
}

var (
	sqlScript = `Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgc3R1ZGVudHMgKAoJaWQgU0VSSUFMIFBSSU1BUlkgS0VZLCBmaXJzdF9uYW1lIFZBUkNIQVIoMjU1KSBOT1QgTlVMTCwgbGFzdF9uYW1lIFZBUkNIQVIoMjU1KSBOT1QgTlVMTCwKCWdlbmRlciBWQVJDSEFSKDUwKSBOT1QgTlVMTCwgZGF0ZV9vZl9iaXJ0aCBEQVRFIE5PVCBOVUxMLCBhZGRyZXNzIFZBUkNIQVIoMjU1KSwKCWNsYXNzIFZBUkNIQVIoMTApIE5PVCBOVUxMLCBzdGF0dXMgVkFSQ0hBUig1MCkgTk9UIE5VTEwgKTsKQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgdGVhY2hlcnMgKAoJaWQgU0VSSUFMIFBSSU1BUlkgS0VZLCBuaXAgVkFSQ0hBUigyNTUpIE5PVCBOVUxMLCBmdWxsbmFtZSBWQVJDSEFSKDI1NSkgTk9UIE5VTEwsCglhZGRyZXNzIFZBUkNIQVIoMjU1KSwgZ3JvdXBzIFZBUkNIQVIoMTApIE5PVCBOVUxMLCBzdGF0dXMgVkFSQ0hBUig1MCkgTk9UIE5VTEwpOwkKSU5TRVJUIElOVE8gdGVhY2hlcnMgKG5pcCwgZnVsbG5hbWUsIGFkZHJlc3MsIGdyb3Vwcywgc3RhdHVzICkgCglWQUxVRVMgKCcxMjM0NTY3ODkwJywgJ0plZmZyZXkgSGFydG9ubycsICdKbC4gSmFsYW4nLCAnQScsICdhY3RpdmUnKSwKCSgnMDk4NzY1NDMyMScsICdIYXJpIFJhaGFydGEnLCAnSmwuIEphbGFuJywgJ0InLCAnYWN0aXZlJyksCgkoJzEyMzQ1Njc4OTEnLCAnSGFybWlvbm8gSnVkaWFudG8nLCAnSmwuIEphbGFuJywgJ0EnLCAnaW5hY3RpdmUnKSwKCSgnMDk4NzY1NDMyMicsICdIbyBTemUgV2FoJywgJ0psLiBKYWxhbicsICdCJywgJ2luYWN0aXZlJyksCgkoJzEyMzQ1Njc4OTInLCAnSmFoamEgU2FudG9zbycsICdKbC4gSmFsYW4nLCAnQScsICdhY3RpdmUnKSwKCSgnMDk4NzY1NDMyMycsICdLdXN1bWF3YXRpJywgJ0psLiBKYWxhbicsICdCJywgJ2FjdGl2ZScpLAoJKCcxMjM0NTY3ODkzJywgJ0x1bmFyZGkgQmFzdWtpJywgJ0psLiBKYWxhbicsICdBJywgJ2FjdGl2ZScpLAoJKCcwOTg3NjU0MzI0JywgJ01vaGFtYWQgTm9lcicsICdKbC4gSmFsYW4nLCAnQycsICdhY3RpdmUnKSwKCSgnMTIzNDU2Nzg5NCcsICdNdWhhbW1hZCBJa2JhbCcsICdKbC4gSmFsYW4nLCAnQycsICdhY3RpdmUnKSwKCSgnMDk4NzY1NDMyNScsICdOaWxhIEZhdGluYScsICdKbC4gSmFsYW4nLCAnQycsICdpbmFjdGl2ZScpOw==`
)

func SQLExecute(db *sql.DB) error {
	sqlScript, _ := base64.StdEncoding.DecodeString(sqlScript)
	_, err := db.Exec(string(sqlScript))
	if err != nil {
		return err
	}

	fmt.Println("success create table and insert data")
	return nil
}

func main() {
	dbCredential := Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "1345678",
		DatabaseName: "my_db",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbConn.Exec("DROP TABLE IF EXISTS students, teachers CASCADE;")
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = InsertSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = UpdateSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteSQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
