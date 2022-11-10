package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query data in table 'people'", Ordered, func() {
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

	When("total data less than equal 5", func() {
		It("should get all data with the heaviest people", func() {
			_, err = dbConn.Exec("DROP TABLE IF EXISTS people CASCADE")
			if err != nil {
				panic(fmt.Sprintf("failed drop table people: %s", err.Error()))
			}

			err = main.SQLExecute(dbConn, main.InsertStr)
			if err != nil {
				panic(err)
			}

			res, err := main.QuerySQL(dbConn)
			if err != nil {
				panic(err)
			}

			Expect(res).To(Equal([]main.People{
				{3, "1234567890125", "Andre William", "1990-01-03", 80, "Jl. Abc"},
				{1, "1234567890123", "Andi Sukirna", "1990-01-01", 70, "Jl. Abc"},
			}))
			Expect(res).To(HaveLen(2))
		})
	})

	When("total data more than 5", func() {
		It("should get top 5 data with the heaviest people", func() {
			_, err = dbConn.Exec("DROP TABLE IF EXISTS people CASCADE")
			if err != nil {
				panic(fmt.Sprintf("failed drop table people %s", err.Error()))
			}

			err = main.SQLExecute(dbConn, `INSERT INTO people VALUES 
			(1, '1234567890123' , 'Andi' , 'Sukirna' , 'laki-laki', '1990-01-01', 170, 70, 'Jl. Jakarta'),
			(2, '1234567890124' , 'Sulis' , 'Indahwati' , 'perempuan', '1990-01-02', 160, 50, 'Jl. Jakarta'),
			(3, '1234567890125' , 'Andre' , 'William' , 'laki-laki', '1990-01-03', 180, 80, 'Jl. Jakarta'),
			(4, '1234567890126' , 'Henny' , 'Welas' , 'perempuan', '1990-01-04', 150, 40, 'Jl. Jakarta'),
			(5, '1234567890127' , 'Wendy' , 'Sukirna' , 'laki-laki', '1990-01-25', 170, 71, 'Jl. Jakarta'),
			(6, '1234567890128' , 'Rendy' , 'Santoso' , 'laki-laki', '1990-01-25', 170, 75, 'Jl. Jakarta'),
			(7, '1234567890129' , 'Rina' , 'Santoso' , 'perempuan', '1990-01-12', 170, 75, 'Jl. Jakarta'),
			(8, '1234567890130' , 'Johan' , 'Roger' , 'laki-laki', '1990-01-10', 170, 69, 'Jl. Jakarta'),
			(9, '1234567890131' , 'Albert' , 'Sunardi' , 'laki-laki', '1990-01-11', 170, 73, 'Jl. Jakarta'),
			(10, '1234567890132' , 'Firman' , 'Hardi' , 'laki-laki', '1990-01-05', 170, 74, 'Jl. Jakarta');`)

			if err != nil {
				panic(err)
			}

			res, err := main.QuerySQL(dbConn)
			if err != nil {
				panic(err)
			}

			Expect(res).To(Equal([]main.People{
				{
					ID:          3,
					NIK:         "1234567890125",
					Fullname:    "Andre William",
					DateOfBirth: "1990-01-03",
					Weight:      80,
					Address:     "Jl. Jakarta",
				},
				{
					ID:          6,
					NIK:         "1234567890128",
					Fullname:    "Rendy Santoso",
					DateOfBirth: "1990-01-25",
					Weight:      75,
					Address:     "Jl. Jakarta",
				},
				{
					ID:          10,
					NIK:         "1234567890132",
					Fullname:    "Firman Hardi",
					DateOfBirth: "1990-01-05",
					Weight:      74,
					Address:     "Jl. Jakarta",
				},
				{
					ID:          9,
					NIK:         "1234567890131",
					Fullname:    "Albert Sunardi",
					DateOfBirth: "1990-01-11",
					Weight:      73,
					Address:     "Jl. Jakarta",
				},
				{
					ID:          5,
					NIK:         "1234567890127",
					Fullname:    "Wendy Sukirna",
					DateOfBirth: "1990-01-25",
					Weight:      71,
					Address:     "Jl. Jakarta",
				},
			}))
			Expect(res).To(HaveLen(5))

		})
	})
})
