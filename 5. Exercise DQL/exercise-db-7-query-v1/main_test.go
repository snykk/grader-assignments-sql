package main_test

import (
	main "a21hc3NpZ25tZW50"

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

	BeforeAll(func() {
		// database/sql version (not using gorm)
		dbConn, err := main.Connect(&dbCredential)
		if err != nil {
			panic("failed connecting to database, please check Connect function: " + err.Error())
		}

		_, err = dbConn.Exec("DROP TABLE IF EXISTS reports CASCADE")
		if err != nil {
			panic("failed dropping table reports: " + err.Error())
		}

		err = main.SQLExecute(dbConn)
		if err != nil {
			panic("failed creating and insert table reports: " + err.Error())
		}
	})

	dbConn, err := main.Connect(&dbCredential)
	if err != nil {
		panic("failed connecting to database, please check Connect function")
	}

	Describe("Query data in table 'reports'", func() {
		It("Should get all data with score less than 70 or absent more than 5", func() {
			res, err := main.QueryStudent(dbConn)
			if err != nil {
				panic("failed inserting data: " + err.Error())
			}

			Expect(res).To(Equal([]main.Report{
				{5, "Andrew Doe", "1A", 60, 6},
				{6, "Rendy Doe", "1B", 69, 2},
				{7, "John Smith", "1B", 69, 6},
				{10, "Wendy Doe", "1B", 40, 7},
			}))
		})
	})

})
