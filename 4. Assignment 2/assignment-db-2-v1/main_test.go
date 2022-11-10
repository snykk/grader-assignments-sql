package main_test

import (
	main "a21hc3NpZ25tZW50"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Student struct {
	Id          int    `sql:"id"`
	FirstName   string `sql:"first_name"`
	LastName    string `sql:"last_name"`
	Gender      string `sql:"gender"`
	DateOfBirth string `sql:"date_of_birth"`
	Address     string `sql:"address"`
	Class       string `sql:"class"`
	Status      string `sql:"status"`
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

	dbConn, err := main.Connect(&dbCredential)
	if err != nil {
		panic("failed connecting to database, please check Connect function")
	}

	Describe("Insert data to table 'students'", func() {
		It("should have length of 15", func() {
			_, err := dbConn.Exec("DROP TABLE IF EXISTS students, teachers CASCADE")
			if err != nil {
				panic("failed to drop table database: " + err.Error())
			}

			err = main.SQLExecute(dbConn)
			if err != nil {
				panic("failed execute SQL: " + err.Error())
			}

			err = main.InsertSQL(dbConn)
			if err != nil {
				panic("failed to insert SQL: " + err.Error())
			}

			// Check length of data
			row := dbConn.QueryRow("SELECT COUNT(*) FROM students")
			if err != nil {
				panic("failed to query data: " + err.Error())
			}

			var count int
			err = row.Scan(&count)
			if err != nil {
				panic("failed to scan data: " + err.Error())
			}
			Expect(count).To(Equal(15))
		})

		It("should insert correct data", func() {
			rows, err := dbConn.Query("SELECT id, first_name, last_name, gender, date_of_birth, COALESCE(address, ''), class, status FROM students")
			if err != nil {
				panic("failed to query data: " + err.Error())
			}

			var students []Student

			for rows.Next() {
				var s Student
				if err = rows.Scan(&s.Id, &s.FirstName, &s.LastName, &s.Gender, &s.DateOfBirth, &s.Address, &s.Class, &s.Status); err != nil {
					panic("failed to scan data: " + err.Error())
				}

				s.DateOfBirth = strings.Split(s.DateOfBirth, "T")[0]
				students = append(students, s)
			}

			expectedStudents := []Student{
				{
					Id:          1,
					FirstName:   "Imam",
					LastName:    "Rendi",
					Gender:      "laki-laki",
					DateOfBirth: "2002-02-02",
					Address:     "Jl Jakarta",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          2,
					FirstName:   "Andi",
					LastName:    "Sukirna",
					Gender:      "laki-laki",
					DateOfBirth: "2002-02-03",
					Address:     "Jl Jakarta",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          3,
					FirstName:   "Achmad",
					LastName:    "Fadjar",
					Gender:      "laki-laki",
					DateOfBirth: "2002-02-03",
					Address:     "Jl Depok",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          4,
					FirstName:   "Achmad",
					LastName:    "Kalla",
					Gender:      "laki-laki",
					DateOfBirth: "2002-02-03",
					Address:     "",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          5,
					FirstName:   "Aida",
					LastName:    "Ishak",
					Gender:      "perempuan",
					DateOfBirth: "2002-01-01",
					Address:     "Jl Depok",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          6,
					FirstName:   "Alice",
					LastName:    "Haryono",
					Gender:      "perempuan",
					DateOfBirth: "2002-01-01",
					Address:     "Jl Depok",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          7,
					FirstName:   "Calvin",
					LastName:    "Lukmantara",
					Gender:      "laki-laki",
					DateOfBirth: "2002-01-05",
					Address:     "Jl Jakarta Utara",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          8,
					FirstName:   "Chris",
					LastName:    "Fong",
					Gender:      "laki-laki",
					DateOfBirth: "2002-01-07",
					Address:     "Jl Jakarta Utara",
					Class:       "1A",
					Status:      "active",
				},
				{
					Id:          9,
					FirstName:   "Citra",
					LastName:    "Andini",
					Gender:      "perempuan",
					DateOfBirth: "2002-01-08",
					Address:     "",
					Class:       "1B",
					Status:      "active",
				},
				{
					Id:          10,
					FirstName:   "Darwin",
					LastName:    "Leo",
					Gender:      "laki-laki",
					DateOfBirth: "2003-04-01",
					Address:     "Jl Jakarta Utara",
					Class:       "1B",
					Status:      "active",
				},
				{
					Id:          11,
					FirstName:   "Dewi",
					LastName:    "Nilka Sari",
					Gender:      "perempuan",
					DateOfBirth: "2003-05-01",
					Address:     "Jl Jakarta Utara",
					Class:       "1B",
					Status:      "active",
				},
				{
					Id:          12,
					FirstName:   "Edy",
					LastName:    "Kosasih",
					Gender:      "laki-laki",
					DateOfBirth: "2003-06-01",
					Address:     "Jl Sukarno Hatta",
					Class:       "1B",
					Status:      "active",
				},
				{
					Id:          13,
					FirstName:   "Fabian",
					LastName:    "Gelael",
					Gender:      "laki-laki",
					DateOfBirth: "2003-07-01",
					Address:     "Jl Sukarno Hatta",
					Class:       "1B",
					Status:      "active",
				},
				{
					Id:          14,
					FirstName:   "Halifah",
					LastName:    "Indah",
					Gender:      "perempuan",
					DateOfBirth: "2003-09-01",
					Address:     "",
					Class:       "1B",
					Status:      "active",
				},
				{
					Id:          15,
					FirstName:   "Hari",
					LastName:    "Widodo",
					Gender:      "laki-laki",
					DateOfBirth: "2003-10-01",
					Address:     "",
					Class:       "1B",
					Status:      "active",
				},
			}

			Expect(len(students)).To(Equal(len(expectedStudents)))
			for i := range students {
				Expect(students[i]).To(Equal(expectedStudents[i]))
			}
		})
	})

	Describe("SQL data table 'teachers'", func() {
		When("Update data 'teachers' groups to B", func() {
			It("should not have data with groups A", func() {
				_, err := dbConn.Exec("DROP TABLE IF EXISTS students, teachers CASCADE")
				if err != nil {
					panic("failed to drop table database: " + err.Error())
				}

				err = main.SQLExecute(dbConn)
				if err != nil {
					panic("failed execute SQL: " + err.Error())
				}

				err = main.UpdateSQL(dbConn)
				if err != nil {
					panic("failed to update SQL: " + err.Error())
				}

				row := dbConn.QueryRow("SELECT COUNT(*) FROM teachers WHERE groups = 'A'")

				var count int
				err = row.Scan(&count)

				Expect(err).To(BeNil())
				Expect(count).To(Equal(0))
			})

			It("add total data teacher with groups B", func() {
				row := dbConn.QueryRow("SELECT COUNT(*) FROM teachers WHERE groups = 'B'")

				var count int
				err = row.Scan(&count)

				Expect(err).To(BeNil())
				Expect(count).To(Equal(7))
			})
		})

		When("Delete data with status 'inactive'", func() {
			It("should not have data with status 'inactive'", func() {
				err = main.DeleteSQL(dbConn)
				if err != nil {
					panic("failed to delete SQL: " + err.Error())
				}

				row := dbConn.QueryRow("SELECT COUNT(*) FROM teachers WHERE status = 'inactive'")

				var count int
				err = row.Scan(&count)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(0))
			})

			It("should get less total data", func() {
				row := dbConn.QueryRow("SELECT COUNT(*) FROM teachers")

				var count int
				err = row.Scan(&count)
				Expect(err).To(BeNil())
				Expect(count).To(Equal(7))
			})
		})
	})
})
