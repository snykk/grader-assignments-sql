package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var expectTable = []main.TableCheck{
	{"id", 1, "NO", "integer", 0},
	{"nik", 2, "NO", "character varying", 255},
	{"fullname", 3, "NO", "character varying", 255},
	{"gender", 4, "NO", "character varying", 50},
	{"birth_date", 5, "NO", "date", 0},
	{"is_married", 6, "YES", "boolean", 0},
	{"height", 7, "YES", "double precision", 0},
	{"weight", 8, "YES", "double precision", 0},
	{"address", 9, "YES", "text", 0},
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
		dbConn, err := main.Connect(&dbCredential)
		if err != nil {
			panic("failed connecting to database, please check Connect function")
		}

		_, err = dbConn.Exec("DROP TABLE IF EXISTS persons CASCADE")
		if err != nil {
			panic("failed dropping table persons")
		}
	})

	dbConn, err := main.Connect(&dbCredential)
	if err != nil {
		panic("failed connecting to database, please check Connect function")
	}

	Describe("Create table 'persons'", func() {
		It("should not error and create table", func() {
			err := main.CreateTable(dbConn)
			Expect(err).To(BeNil())

			row := dbConn.QueryRow(`SELECT listTable.* FROM (
				SELECT
					table_name as tables
				FROM information_schema.tables
				WHERE table_type = 'BASE TABLE' 
				AND table_schema NOT IN ('pg_catalog', 'information_schema')
			) listTable
			WHERE listTable.tables = 'persons';`)

			var table string
			err = row.Scan(&table)
			Expect(err).To(BeNil())

			Expect(table).To(Equal("persons"))
		})

		It("should have 9 columns with exact constrainsts", func() {
			rows, err := dbConn.Query(`SELECT 
				column_name,
				ordinal_position,
				is_nullable,
				data_type,
				coalesce(character_maximum_length, 0) as character_maximum_length
			FROM INFORMATION_SCHEMA.COLUMNS where TABLE_NAME='persons'
			ORDER BY ordinal_position ASC
			`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			var tableChecks []main.TableCheck
			for rows.Next() {
				var tableCheck main.TableCheck
				err := rows.Scan(&tableCheck.ColumnName, &tableCheck.OrdinalPosition, &tableCheck.IsNullable, &tableCheck.DataType, &tableCheck.CharLength)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				tableChecks = append(tableChecks, tableCheck)
			}

			for i := 0; i < len(tableChecks); i++ {
				Expect(tableChecks[i]).To(Equal(expectTable[i]))
			}
		})

		It("should have primary key and unique data", func() {
			var constraints []struct {
				ConstraintName string `sql:"constraint_name"`
				ColumnName     string `sql:"column_name"`
			}

			rows, err := dbConn.Query(`SELECT constraint_type, column_name FROM information_schema.table_constraints JOIN information_schema.constraint_column_usage USING (constraint_schema, constraint_name) WHERE information_schema.
			constraint_column_usage.table_name='persons';`)

			if err != nil {
				panic(fmt.Sprintf("failed query table: %s", err.Error()))
			}

			for rows.Next() {
				var constraint struct {
					ConstraintName string `sql:"constraint_name"`
					ColumnName     string `sql:"column_name"`
				}

				err := rows.Scan(&constraint.ConstraintName, &constraint.ColumnName)
				if err != nil {
					panic(fmt.Sprintf("failed scan row: %s", err.Error()))
				}

				constraints = append(constraints, constraint)
			}

			Expect(constraints[0].ConstraintName).To(Equal("PRIMARY KEY"))
			Expect(constraints[0].ColumnName).To(Equal("id"))

			Expect(constraints[1].ConstraintName).To(Equal("UNIQUE"))
			Expect(constraints[1].ColumnName).To(Equal("nik"))
		})
	})
})
