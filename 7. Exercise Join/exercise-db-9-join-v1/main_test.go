package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", Ordered, func() {
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

	Describe("Query join students and reports", func() {
		_, err := dbConn.Exec("DROP TABLE IF EXISTS students, reports CASCADE")
		if err != nil {
			panic(fmt.Sprintf("failed drop table students & reports: %s", err.Error()))
		}

		err = main.SQLExecute(dbConn, main.InsertSQL)
		if err != nil {
			panic(fmt.Sprintf("failed inserting data: %s", err.Error()))
		}

		When("data meet the criteria", func() {
			It("should return the correct result", func() {

				res, err := main.QueryJoinSQL(dbConn)
				if err != nil {
					panic(fmt.Sprintf("failed querying data: %s", err.Error()))
				}

				Expect(res).To(Equal([]main.Report{
					{
						ID:       14,
						Fullname: "Bob Abrams",
						Class:    "1B",
						Status:   "active",
						Study:    "English",
						Score:    30,
					},
					{
						ID:       15,
						Fullname: "Bob Abrams",
						Class:    "1B",
						Status:   "active",
						Study:    "Science",
						Score:    40,
					},
					{
						ID:       16,
						Fullname: "Bob Abrams",
						Class:    "1B",
						Status:   "active",
						Study:    "Indonesia",
						Score:    50,
					},
					{
						ID:       5,
						Fullname: "Jane Willy",
						Class:    "1A",
						Status:   "active",
						Study:    "Math",
						Score:    55,
					},
					{
						ID:       7,
						Fullname: "Jane Willy",
						Class:    "1A",
						Status:   "active",
						Study:    "Science",
						Score:    61,
					},
					{
						ID:       13,
						Fullname: "Bob Abrams",
						Class:    "1B",
						Status:   "active",
						Study:    "Math",
						Score:    65,
					},
				}))
			})
		})
	})
})
