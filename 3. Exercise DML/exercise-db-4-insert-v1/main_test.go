package main_test

import (
	main "a21hc3NpZ25tZW50"
	"log"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Students struct {
	ID          int    `sql:"id"`
	FirstName   string `sql:"first_name"`
	LastName    string `sql:"last_name"`
	DateOfBirth string `sql:"date_of_birth"`
	Address     string `sql:"address"`
	Class       string `sql:"class"`
	Status      string `sql:"status"`
}

var Expected = []Students{
	{1, "Abdi", "Doe", "2003-12-01", "Jakarta", "1A", "active"},
	{2, "Jane", "Doe", "2004-02-01", "Jakarta", "1A", "active"},
	{3, "Bernard", "Smith", "2004-02-01", "Jakarta", "1A", "active"},
	{4, "Jane", "Smith", "2003-12-02", "Jakarta", "1B", "active"},
	{5, "Andrew", "Doe", "2004-07-04", "Jakarta", "1B", "inactive"},
	{6, "Rendy", "Doe", "2004-06-10", "Jakarta", "1B", "inactive"},
	{7, "John", "Smith", "2004-05-11", "Jakarta", "1B", "inactive"},
	{8, "Herry", "Smith", "2004-04-12", "", "1B", "active"},
	{9, "John", "William", "2004-03-20", "", "1B", "active"},
	{10, "Wendy", "Doe", "2004-02-21", "", "1B", "active"},
}

func ChangeToDateStr(date string) string {
	return strings.Split(date, "T")[0]
}

var _ = Describe("Main", Ordered, func() {
	dbCredential := main.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
		Schema:       "public",
	}
	BeforeAll(func() {
		// database/sql version (not using gorm)
		dbConn, err := main.Connect(&dbCredential)
		if err != nil {
			panic("failed connecting to database, please check Connect function: " + err.Error())
		}

		_, err = dbConn.Exec("DROP TABLE IF EXISTS students CASCADE")
		if err != nil {
			panic("failed dropping table students: " + err.Error())
		}

		err = main.CreateTable(dbConn)
		if err != nil {
			panic("failed creating table students: " + err.Error())
		}
	})

	dbConn, err := main.Connect(&dbCredential)
	if err != nil {
		panic("failed connecting to database, please check Connect function")
	}

	Describe("Insert 10 data to table 'students'", func() {
		It("should have 10 rows", func() {
			err = main.InsertSQL(dbConn)
			if err != nil {
				panic("failed inserting data")
			}

			var count int
			err = dbConn.QueryRow("SELECT COUNT(*) FROM students").Scan(&count)
			if err != nil {
				panic("failed counting data")
			}

			Expect(count).To(Equal(10))
		})

		It("should same as data table", func() {
			rows, err := dbConn.Query(`SELECT 
			id, first_name, last_name, date_of_birth, COALESCE(address, ''), class, status 
			FROM students ORDER BY id ASC`)
			if err != nil {
				panic("failed querying data")
			}

			var results []Students

			for rows.Next() {
				var res Students

				err = rows.Scan(&res.ID, &res.FirstName, &res.LastName, &res.DateOfBirth, &res.Address, &res.Class, &res.Status)
				if err != nil {
					log.Println(err.Error())
					panic("failed scanning data")
				}

				res.DateOfBirth = ChangeToDateStr(res.DateOfBirth)
				results = append(results, res)
			}

			for i, res := range results {
				Expect(res).To(Equal(Expected[i]))
			}
		})
	})
})
