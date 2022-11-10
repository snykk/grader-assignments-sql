package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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

type People struct {
	ID          int    `sql:"id"`
	NIK         string `sql:"nik"`
	Fullname    string `sql:"fullname"`
	DateOfBirth string `sql:"date_of_birth"`
	Weight      int    `sql:"weight"`
	Address     string `sql:"address"`
}

func ChangeToDateStr(date string) string {
	return strings.Split(date, "T")[0]
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

func QuerySQL(db *sql.DB) ([]People, error) {
	var res []People

	rows, err := db.Query(queryStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p People

		if err := rows.Scan(&p.ID, &p.NIK, &p.Fullname, &p.DateOfBirth, &p.Weight, &p.Address); err != nil {
			log.Fatal(err)
		}

		p.DateOfBirth = ChangeToDateStr(p.DateOfBirth)
		res = append(res, p)
	}

	return res, nil
}

var (
	CreateStr = `CREATE TABLE IF NOT EXISTS people (
		id INT PRIMARY KEY,
		NIK VARCHAR(255) NOT NULL,
		first_name VARCHAR(150) NOT NULL,
		last_name VARCHAR(150) NOT NULL,
		gender VARCHAR(50) NOT NULL,
		date_of_birth DATE NOT NULL,
		height INT NOT NULL,
		weight INT NOT NULL,
		address VARCHAR(255)
	)`

	InsertStr = `INSERT INTO people VALUES (1, '1234567890123' , 'Andi' , 'Sukirna' , 'laki-laki', '1990-01-01', 170, 70, 'Jl. Abc'),
	(2, '1234567890124' , 'Sulis' , 'Indahwati' , 'perempuan', '1990-01-02', 160, 50, 'Jl. Abc'),
	(3, '1234567890125' , 'Andre' , 'William' , 'laki-laki', '1990-01-03', 180, 80, 'Jl. Abc'),
	(4, '1234567890126' , 'Henny' , 'Welas' , 'perempuan', '1990-01-04', 150, 40, 'Jl. Abc');`
)

func SQLExecute(db *sql.DB, InsertStr string) error {
	_, err := db.Exec(CreateStr)
	if err != nil {
		return err
	}

	fmt.Println("success create table")

	_, err = db.Exec(InsertStr)
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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS peoples CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn, InsertStr)
	if err != nil {
		log.Fatal(err)
	}

	res, err := QuerySQL(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range res {
		fmt.Println(s)
	}
}
