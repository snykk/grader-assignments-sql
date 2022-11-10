package main_test

import (
	main "a21hc3NpZ25tZW51"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	//Change this with your database credential
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

	When("add new column to table 'students'", func() {

		_, err = dbConn.Exec("DROP TABLE IF EXISTS students CASCADE")
		if err != nil {
			panic("failed dropping table students")
		}

		err = main.CreateTable(dbConn)
		if err != nil {
			panic("failed creating table students")
		}

		It("should have new 5 column", func() {
			err := main.AlterAdd(dbConn)
			Expect(err).To(BeNil())

			rows, err := dbConn.Query(`SELECT 
					column_name,
					is_nullable,
					data_type
				FROM INFORMATION_SCHEMA.COLUMNS 
				where TABLE_NAME='students' 
					and column_name in ('date_of_birth', 'street', 'city', 'province', 'country', 'postal_code')
				ORDER BY ordinal_position ASC`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			var columnsCheck []main.Column
			for rows.Next() {
				var colCheck main.Column
				err := rows.Scan(&colCheck.ColumnName, &colCheck.IsNullable, &colCheck.DataType)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				columnsCheck = append(columnsCheck, colCheck)
			}

			var expected []main.Column = []main.Column{
				{"date_of_birth", "NO", "date"},
				{"street", "YES", "character varying"},
				{"city", "YES", "character varying"},
				{"province", "YES", "character varying"},
				{"country", "YES", "character varying"},
				{"postal_code", "YES", "character varying"},
			}

			for i := 0; i < len(expected); i++ {
				Expect(columnsCheck).To(ContainElement(expected[i]))
			}

			Expect(columnsCheck).To(HaveLen(6))
		})
	})

	When("drop column table 'students'", func() {
		_, err = dbConn.Exec("DROP TABLE IF EXISTS students CASCADE")
		if err != nil {
			panic("failed dropping table students")
		}

		err = main.CreateTable(dbConn)
		if err != nil {
			panic("failed creating table students")
		}

		It("should remove 4 column", func() {
			err := main.AlterDrop(dbConn)
			Expect(err).To(BeNil())

			rows, err := dbConn.Query(`SELECT 
					column_name,
					is_nullable,
					data_type
				FROM INFORMATION_SCHEMA.COLUMNS 
				where TABLE_NAME='students' 
					and column_name in ('address', 'day_of_birth', 'month_of_birth', 'year_of_birth')
				ORDER BY ordinal_position ASC`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			var columnsCheck []main.Column
			for rows.Next() {
				var colCheck main.Column
				err := rows.Scan(&colCheck.ColumnName, &colCheck.IsNullable, &colCheck.DataType)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				columnsCheck = append(columnsCheck, colCheck)
			}

			Expect(columnsCheck).To(BeEmpty())
			Expect(columnsCheck).To(HaveLen(0))
		})
	})

	When("modify column table 'students'", func() {
		_, err = dbConn.Exec("DROP TABLE IF EXISTS students CASCADE")
		if err != nil {
			panic("failed dropping table students")
		}

		err = main.CreateTable(dbConn)
		if err != nil {
			panic("failed creating table students")
		}

		It("should change the type of grade", func() {
			err := main.AlterModify(dbConn)
			Expect(err).To(BeNil())

			rows, err := dbConn.Query(`SELECT 
					column_name,
					is_nullable,
					data_type
				FROM INFORMATION_SCHEMA.COLUMNS 
				where TABLE_NAME='students'
					and column_name in ('grade')
				ORDER BY ordinal_position ASC`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			var columnsCheck []main.Column
			for rows.Next() {
				var colCheck main.Column
				err := rows.Scan(&colCheck.ColumnName, &colCheck.IsNullable, &colCheck.DataType)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				columnsCheck = append(columnsCheck, colCheck)
			}

			Expect(columnsCheck).To(Equal([]main.Column{
				{"grade", "YES", "double precision"},
			}))
			Expect(columnsCheck).To(HaveLen(1))
		})
	})
})
