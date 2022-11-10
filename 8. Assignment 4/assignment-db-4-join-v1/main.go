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

type Order struct {
	OrderID     int    `sql:"order_id"`
	Fullname    string `sql:"fullname"`
	Email       string `sql:"email"`
	ProductName string `sql:"product_name"`
	UnitPrice   int    `sql:"unit_price"`
	Quantity    int    `sql:"quantity"`
	OrderDate   string `sql:"order_date"`
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

//go:embed join.sql
var joinStr string

func QueryJoinSQL(db *sql.DB) ([]Order, error) {
	var res []Order

	rows, err := db.Query(joinStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var o Order

		if err := rows.Scan(&o.OrderID, &o.Fullname, &o.Email, &o.ProductName, &o.UnitPrice, &o.Quantity, &o.OrderDate); err != nil {
			return nil, err
		}

		o.OrderDate = ChangeToDateStr(o.OrderDate)
		res = append(res, o)
	}

	return res, nil
}

var (
	CreateStr = `CREATE table IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, fullname VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL UNIQUE,
		address VARCHAR(255) NOT NULL, status VARCHAR(255) NOT NULL);
	
	CREATE table IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY, user_id INT NOT NULL, product_name VARCHAR(255) NOT NULL,
		unit_price INT NOT NULL, quantity INT NOT NULL, order_date DATE NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id));`

	InserStr = `INSERT INTO users (fullname, email, address, status) 
	VALUES ('John Doe', 'john@mail.com', 'Jl. Kebon Jeruk', 'active'),
	('Jane Doe', 'jane@mail.com', 'Jl. Kebon Jeruk', 'active'),
	('Bob Doe', 'bobdoe@mail.com', 'Jl. Kebon Jeruk', 'active'),
	('Alice Doe', 'alice@mail.com', 'Jl. Kebon Jeruk', 'active'),
	('Bob Marley', 'marleybob@mail.com', 'Jl. Kebon Jeruk', 'inactive');
	
	INSERT INTO orders (user_id, product_name, unit_price, quantity, order_date)
	VALUES (1, 'Beras 3kg', 30000, 10, '2021-01-01'),
	(2, 'Gula 2kg', 20000, 5, '2021-01-01'),
	(3, 'Beras 10kg', 100000, 6, '2021-01-01'),
	(4, 'Telur', 5000, 50, '2021-01-01'),
	(5, 'Minyak Goreng 1Lt', 30000, 17, '2021-01-01');
	`
)

func SQLExecute(db *sql.DB, insertSQL string) error {
	_, err := db.Exec(CreateStr)
	if err != nil {
		return err
	}

	fmt.Println("success create table")

	_, err = db.Exec(insertSQL)
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

	_, err = dbConn.Exec("DROP TABLE IF EXISTS users, orders CASCADE")
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn, InserStr)
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
