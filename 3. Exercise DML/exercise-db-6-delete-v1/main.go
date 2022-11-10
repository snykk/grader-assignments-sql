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

//go:embed delete.sql
var deleteSQL string

func DeleteDataStudents(dbConn *sql.DB) error {
	_, err := dbConn.Exec(deleteSQL)
	if err != nil {
		return err
	}

	fmt.Println("success delete data")
	return nil
}

var (
	sqlScript1 = "Q1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgc3R1ZGVudHMgKAoJCWlkIElOVCwgCgkJZmlyc3RfbmFtZSBWQVJDSEFSKDEwMCksCgkJbGFzdF9uYW1lIFZBUkNIQVIoMTAwKSwKCQlkYXRlX29mX2JpcnRoIERBVEUsCgkJYWRkcmVzcyBWQVJDSEFSKDI1NSksCgkJY2xhc3MgVkFSQ0hBUigxMDApLAoJCXN0YXR1cyBWQVJDSEFSKDEwMCkKCSk="
	sqlScript2 = "SU5TRVJUIElOVE8gc3R1ZGVudHMgKGlkLCBmaXJzdF9uYW1lLCBsYXN0X25hbWUsIGRhdGVfb2ZfYmlydGgsIGFkZHJlc3MsIGNsYXNzLCBzdGF0dXMpIApWQUxVRVMgKDEsICdBYmRpJywgJ0RvZScsICcyMDAzLTEyLTAxJywgJ0pha2FydGEnLCAnMUEnLCAnYWN0aXZlJyksCigyLCAnSmFuZScsICdEb2UnLCAnMjAwNC0wMi0wMScsICdKYWthcnRhJywgJzFBJywgJ2FjdGl2ZScpLAooMywgJ0Jlcm5hcmQnLCAnU21pdGgnLCAnMjAwNC0wMi0wMScsICdKYWthcnRhJywgJzFBJywgJ2FjdGl2ZScpLAooNCwgJ0phbmUnLCAnU21pdGgnLCAnMjAwMy0xMi0wMicsICdKYWthcnRhJywgJzFCJywgJ2FjdGl2ZScpLAooNSwgJ0FuZHJldycsICdEb2UnLCAnMjAwNC0wNy0wNCcsICdKYWthcnRhJywgJzFCJywgJ2luYWN0aXZlJyksCig2LCAnUmVuZHknLCAnRG9lJywgJzIwMDQtMDYtMTAnLCAnSmFrYXJ0YScsICcxQicsICdpbmFjdGl2ZScpLAooNywgJ0pvaG4nLCAnU21pdGgnLCAnMjAwNC0wNS0xMScsICdKYWthcnRhJywgJzFCJywgJ2luYWN0aXZlJyksCig4LCAnSGVycnknLCAnU21pdGgnLCAnMjAwNC0wNC0xMicsICdKYWthcnRhJywgJzFCJywgJ2FjdGl2ZScpLAooOSwgJ0pvaG4nLCAnV2lsbGlhbScsICcyMDA0LTAzLTIwJywnSmFrYXJ0YScsICcxQicsICdhY3RpdmUnKSwKKDEwLCAnV2VuZHknLCAnRG9lJywgJzIwMDQtMDItMjEnLCAnSmFrYXJ0YScsICcxQicsICdhY3RpdmUnKTs="
)

func SQLExecute(dbConn *sql.DB) error {
	sqlScript1, _ := base64.StdEncoding.DecodeString(sqlScript1)
	_, err := dbConn.Exec(string(sqlScript1))

	if err != nil {
		return err
	}

	fmt.Println("success create table")

	sqlScript2, _ := base64.StdEncoding.DecodeString(sqlScript2)
	_, err = dbConn.Exec(string(sqlScript2))
	if err != nil {
		return err
	}

	fmt.Println("success insert data")

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

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteDataStudents(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}
