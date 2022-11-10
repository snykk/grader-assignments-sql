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

	Describe("alter table 'users'", func() {
		dbConn, err := main.Connect(&dbCredential)
		if err != nil {
			panic("failed connecting to database, please check Connect function")
		}

		_, err = dbConn.Exec("DROP TABLE IF EXISTS users, attendances, presences CASCADE")
		if err != nil {
			panic(fmt.Sprintf("failed drop table database: %s", err.Error()))
		}

		It("should not error and alter table with exact column", func() {
			err := main.SQLExecute(dbConn)
			if err != nil {
				panic(fmt.Sprintf("failed execute: %s", err.Error()))
			}

			err = main.AlterSQL(dbConn)
			if err != nil {
				panic(fmt.Sprintf("failed alter table: %s", err.Error()))
			}

			rows, err := dbConn.Query(`SELECT 
				column_name,
				ordinal_position,
				is_nullable,
				data_type
			FROM INFORMATION_SCHEMA.COLUMNS where TABLE_NAME='users'
			ORDER BY ordinal_position ASC
			`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			var tableChecks []main.TableCheck
			for rows.Next() {
				var tableCheck main.TableCheck
				err := rows.Scan(&tableCheck.ColumnName, &tableCheck.OrdinalPosition, &tableCheck.IsNullable, &tableCheck.DataType)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				tableChecks = append(tableChecks, tableCheck)
			}

			var expected = []main.TableCheck{
				{ColumnName: "id", OrdinalPosition: 1, IsNullable: "NO", DataType: "integer"},
				{
					ColumnName:      "fullname",
					OrdinalPosition: 2,
					IsNullable:      "NO",
					DataType:        "character varying",
				},
				{
					ColumnName:      "email",
					OrdinalPosition: 3,
					IsNullable:      "NO",
					DataType:        "character varying",
				},
				{
					ColumnName:      "password",
					OrdinalPosition: 4,
					IsNullable:      "NO",
					DataType:        "character varying",
				},
				{
					ColumnName:      "phone",
					OrdinalPosition: 6,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "address",
					OrdinalPosition: 7,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "department",
					OrdinalPosition: 8,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "division",
					OrdinalPosition: 9,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "position",
					OrdinalPosition: 10,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
			}

			for i := 0; i < len(expected); i++ {
				Expect(tableChecks).To(ContainElement(expected[i]))
			}
		})
	})

	Describe("Drop table 'attendances'", func() {
		dbConn, err := main.Connect(&dbCredential)
		if err != nil {
			panic("failed connecting to database, please check Connect function")
		}

		_, err = dbConn.Exec("DROP TABLE IF EXISTS users, attendances, presences CASCADE")
		if err != nil {
			panic(fmt.Sprintf("failed drop table database: %s", err.Error()))
		}

		It("should not error and drop table", func() {
			err := main.SQLExecute(dbConn)
			if err != nil {
				panic(fmt.Sprintf("failed execute: %s", err.Error()))
			}

			err = main.DropSQL(dbConn)
			if err != nil {
				panic(fmt.Sprintf("failed drop table: %s", err.Error()))
			}

			row := dbConn.QueryRow(`SELECT listTable.* FROM (
				SELECT
					table_name as tables
				FROM information_schema.tables
				WHERE table_type = 'BASE TABLE' 
				AND table_schema NOT IN ('pg_catalog', 'information_schema')
			) listTable
			WHERE listTable.tables = 'attendances';`)

			var tableName string
			err = row.Scan(&tableName)
			Expect(err).ToNot(BeNil())

			Expect(tableName).To(Equal(""))
		})
	})

	Describe("Create table 'presences'", func() {
		dbConn, err := main.Connect(&dbCredential)
		if err != nil {
			panic("failed connecting to database, please check Connect function")
		}

		_, err = dbConn.Exec("DROP TABLE IF EXISTS users, attendances, presences CASCADE")
		if err != nil {
			panic(fmt.Sprintf("failed drop table database: %s", err.Error()))
		}

		It("should not error and create table", func() {
			err := main.CreateSQL(dbConn)
			Expect(err).To(BeNil())

			row := dbConn.QueryRow(`SELECT listTable.* FROM (
				SELECT
					table_name as tables
				FROM information_schema.tables
				WHERE table_type = 'BASE TABLE' 
				AND table_schema NOT IN ('pg_catalog', 'information_schema')
			) listTable
			WHERE listTable.tables = 'presences';`)

			var tableName string
			err = row.Scan(&tableName)
			Expect(err).To(BeNil())

			Expect(tableName).To(Equal("presences"))
		})

		It("should have 8 columns with exact constrainsts", func() {
			rows, err := dbConn.Query(`SELECT 
				column_name,
				ordinal_position,
				is_nullable,
				data_type
			FROM INFORMATION_SCHEMA.COLUMNS where TABLE_NAME='presences'
			ORDER BY ordinal_position ASC
			`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			var tableChecks []main.TableCheck
			for rows.Next() {
				var tableCheck main.TableCheck
				err := rows.Scan(&tableCheck.ColumnName, &tableCheck.OrdinalPosition, &tableCheck.IsNullable, &tableCheck.DataType)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				tableChecks = append(tableChecks, tableCheck)
			}

			var expected = []main.TableCheck{
				{ColumnName: "id", OrdinalPosition: 1, IsNullable: "NO", DataType: "integer"},
				{ColumnName: "user_id", OrdinalPosition: 2,
					IsNullable: "NO",
					DataType:   "integer",
				},
				{
					ColumnName:      "presence_date",
					OrdinalPosition: 3,
					IsNullable:      "NO",
					DataType:        "date",
				},
				{
					ColumnName:      "status",
					OrdinalPosition: 4,
					IsNullable:      "NO",
					DataType:        "character varying",
				},
				{
					ColumnName:      "location",
					OrdinalPosition: 5,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "description",
					OrdinalPosition: 6,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "image_presence",
					OrdinalPosition: 7,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
				{
					ColumnName:      "image_location",
					OrdinalPosition: 8,
					IsNullable:      "YES",
					DataType:        "character varying",
				},
			}

			for i := 0; i < len(expected); i++ {
				Expect(tableChecks).To(ContainElement(expected[i]))
			}
		})
	})
})
