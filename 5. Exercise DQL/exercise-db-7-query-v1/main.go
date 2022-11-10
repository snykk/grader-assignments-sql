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

func Connect(creds *Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using database/sql + pq
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

type Report struct {
	Id           int    `sql:"id"`
	StudentName  string `sql:"student_name"`
	StudentClass string `sql:"student_class"`
	FinalScore   int    `sql:"final_score"`
	Absent       int    `sql:"absent"`
}

//go:embed select.sql
var queryStr string

func QueryStudent(db *sql.DB) ([]Report, error) {
	var res []Report

	rows, err := db.Query(queryStr)
	if err != nil {
		log.Println("error query", err)
		return nil, err
	}

	for rows.Next() {
		var currentStr Report
		err = rows.Scan(&currentStr.Id, &currentStr.StudentName, &currentStr.StudentClass, &currentStr.FinalScore, &currentStr.Absent)
		if err != nil {
			log.Println("error scan", err)
			return nil, err
		}

		res = append(res, currentStr)
	}

	return res, nil
}

var (
	sqlScript1 = `CREATE TABLE IF NOT EXISTS reports (
		id INT PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		student_class VARCHAR(100),
		final_score INT NOT NULL,
		absent INT NOT NULL
	)`

	sqlScript2 = `INSERT INTO reports (id, first_name, last_name, student_class, final_score, absent) VALUES
		(1, 'Abdi', 'Doe', '1A', 80, 0),
		(2, 'Jane', 'Doe', '1A', 95, 0),
		(3, 'Bernard', 'Smith', '1A', 95, 0),
		(4, 'Jane', 'Smith', '1A', 86, 4),
		(5, 'Andrew', 'Doe', '1A', 60, 6),
		(6, 'Rendy', 'Doe', '1B', 69, 2),
		(7, 'John', 'Smith', '1B', 69, 6),
		(8, 'Herry', 'Smith', '1B', 91, 3),
		(9, 'John', 'William', '1B', 94, 2),
		(10, 'Wendy', 'Doe', '1B', 40, 7);
	`
)

func SQLExecute(dbConn *sql.DB) error {
	_, err := dbConn.Exec(sqlScript1)

	if err != nil {
		return err
	}

	fmt.Println("success create table")
	_, err = dbConn.Exec(sqlScript2)
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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS raports CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	res, err := QueryStudent(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range res {
		fmt.Println(s)
	}
}
