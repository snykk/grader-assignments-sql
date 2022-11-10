package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query join table 'orders'", Ordered, func() {
	dbCredential := main.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
		Schema:       "public",
	}

	dbConn, err := main.Connect(&dbCredential)
	if err != nil {
		panic("failed connecting to database, please check Connect function")
	}

	When("data meet the criteria", func() {
		It("should return the correct result", func() {
			_, err = dbConn.Exec("DROP TABLE IF EXISTS users, orders CASCADE")
			if err != nil {
				panic(fmt.Sprintf("failed drop table users & orderd: %s", err.Error()))
			}

			err := main.SQLExecute(dbConn, main.InserStr)
			if err != nil {
				panic(fmt.Sprintf("failed inserting data: %s", err.Error()))
			}

			res, err := main.QueryJoinSQL(dbConn)
			if err != nil {
				panic(fmt.Sprintf("failed querying data: %s", err.Error()))
			}

			Expect(res).To(Equal([]main.Order{
				{
					OrderID:     3,
					Fullname:    "Bob Doe",
					Email:       "bobdoe@mail.com",
					ProductName: "Beras 10kg",
					UnitPrice:   100000,
					Quantity:    6,
					OrderDate:   "2021-01-01",
				},
				{
					OrderID:     4,
					Fullname:    "Alice Doe",
					Email:       "alice@mail.com",
					ProductName: "Telur",
					UnitPrice:   5000,
					Quantity:    50,
					OrderDate:   "2021-01-01",
				},
			}))
			Expect(len(res)).To(Equal(2))
		})
	})

	When("only 1 data meet the criteria", func() {
		It("should return 1 correct result", func() {
			_, err = dbConn.Exec("DROP TABLE IF EXISTS users, orders CASCADE")
			if err != nil {
				panic(fmt.Sprintf("failed drop table users & orderd: %s", err.Error()))
			}

			err := main.SQLExecute(dbConn, `
				INSERT INTO users (fullname, email, address, status) 
				VALUES ('John Doe', 'john@mail.com', 'Jl. Kebon Jeruk', 'active'),
				('Jane Doe', 'jane@mail.com', 'Jl. Kebon Jeruk', 'active'),
				('Bob Doe', 'bobdoe@mail.com', 'Jl. Kebon Jeruk', 'inactive'),
				('Alice Doe', 'alice@mail.com', 'Jl. Kebon Jeruk', 'active'),
				('Bob Marley', 'marleybob@mail.com', 'Jl. Kebon Jeruk', 'active');
				
				INSERT INTO orders (user_id, product_name, unit_price, quantity, order_date)
				VALUES (1, 'Beras 3kg', 30000, 10, '2021-01-01'),
				(2, 'Gula 2kg', 20000, 5, '2021-01-01'),
				(4, 'Telur', 5000, 10, '2021-01-01'),
				(5, 'Minyak Goreng 1Lt', 30000, 17, '2021-01-01');
			`)
			if err != nil {
				panic(fmt.Sprintf("failed inserting data: %s", err.Error()))
			}

			res, err := main.QueryJoinSQL(dbConn)
			if err != nil {
				panic(fmt.Sprintf("failed querying data: %s", err.Error()))
			}

			Expect(res).To(Equal([]main.Order{{
				OrderID:     4,
				Fullname:    "Bob Marley",
				Email:       "marleybob@mail.com",
				ProductName: "Minyak Goreng 1Lt",
				UnitPrice:   30000,
				Quantity:    17,
				OrderDate:   "2021-01-01",
			}}))
			Expect(len(res)).To(Equal(1))
		})
	})
})
