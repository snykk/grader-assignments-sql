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

type FinalScore struct {
	ID           int     `sql:"id"`
	Fullname     string  `sql:"fullname"`
	Class        string  `sql:"class"`
	AverageScore float64 `sql:"average_score"`
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

//go:embed select.sql
var queryStr string

func QuerySQL(db *sql.DB) ([]FinalScore, error) {
	var res []FinalScore

	rows, err := db.Query(queryStr)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var f FinalScore

		if err := rows.Scan(&f.ID, &f.Fullname, &f.Class, &f.AverageScore); err != nil {
			return res, err
		}

		res = append(res, f)
	}

	return res, nil
}

var (
	CreateSQL = `CREATE TABLE IF NOT EXISTS final_scores (
	id SERIAL PRIMARY KEY,
	exam_id VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	bahasa_indonesia INT NOT NULL,
	bahasa_inggris INT NOT NULL,
	matematika INT NOT NULL,
	ipa INT NOT NULL,
	exam_status VARCHAR(50) NOT NULL,
	fee_status VARCHAR(50) NOT NULL );`

	InsertSQL = `INSERT INTO final_scores (exam_id, first_name, last_name, bahasa_indonesia, bahasa_inggris, matematika, ipa, exam_status, fee_status) VALUES
	('1A-001', 'John', 'Doe', 80, 90, 70, 80, 'pass', 'full'),
	('1A-002', 'Jane', 'Doe', 90, 80, 90, 80, 'pass', 'full'),
	('1B-003', 'John', 'Smith', 70, 80, 70, 80, 'pass', 'full'),
	('1B-004', 'Jane', 'White', 80, 70, 80, 80, 'pass', 'full'),
	('1B-005', 'John', 'Bernard', 80, 90, 70, 80, 'pass', 'full'),
	('1B-006', 'Jane', 'Abrams', 90, 80, 90, 80, 'pass', 'full'),
	('1B-007', 'John', 'Albert', 70, 80, 70, 80, 'pass', 'full');
	`
)

func SQLExecute(db *sql.DB, insert string) error {
	_, err := db.Exec(CreateSQL)
	if err != nil {
		return err
	}

	fmt.Println("success create table")

	_, err = db.Exec(insert)
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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS final_scores CASCADE")
	if err != nil {
		log.Fatal("error drop table: " + err.Error())
	}

	err = SQLExecute(dbConn, InsertSQL)
	if err != nil {
		log.Fatal("error SQL execute: " + err.Error())
	}

	res, err := QuerySQL(dbConn)
	if err != nil {
		log.Fatal("query error: " + err.Error())
	}

	for _, s := range res {
		fmt.Println(s)
	}
}
