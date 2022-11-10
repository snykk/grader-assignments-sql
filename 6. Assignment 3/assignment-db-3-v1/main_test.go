package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query data table 'final_score'", Ordered, func() {
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

	When("all data column 'exam_status' contain 'pass' and 'fee_status' contain 'full'", func() {
		It("get top 5 student with highest average socre", func() {
			_, err := dbConn.Exec("DROP TABLE IF EXISTS final_scores CASCADE")
			if err != nil {
				panic("error drop table: " + err.Error())
			}

			err = main.SQLExecute(dbConn, main.InsertSQL)
			if err != nil {
				panic("error SQL execute: " + err.Error())
			}

			res, err := main.QuerySQL(dbConn)
			Expect(err).To(BeNil())

			Expect(res).To(Equal([]main.FinalScore{
				{ID: 2, Fullname: "Jane Doe", Class: "1A", AverageScore: 85},
				{ID: 6, Fullname: "Jane Abrams", Class: "1B", AverageScore: 85},
				{ID: 5, Fullname: "John Bernard", Class: "1B", AverageScore: 80},
				{ID: 1, Fullname: "John Doe", Class: "1A", AverageScore: 80},
				{ID: 4, Fullname: "Jane White", Class: "1B", AverageScore: 77},
			}))

			Expect(len(res)).To(Equal(5))
		})
	})

	When("all data column 'exam_status' contain all condition", func() {
		It("get top 5 student with highest average score", func() {
			_, err := dbConn.Exec("DROP TABLE IF EXISTS final_scores CASCADE")
			if err != nil {
				panic("error drop table: " + err.Error())
			}

			err = main.SQLExecute(dbConn, `INSERT INTO final_scores (exam_id, first_name, last_name, bahasa_indonesia, bahasa_inggris, matematika, ipa, exam_status, fee_status) 
			VALUES ('1A-001', 'John', 'Doe', 80, 90, 70, 80, 'pass', 'full'),
			('1A-002', 'Jane', 'Doe', 90, 80, 90, 80, 'pass', 'not paid'),
			('1A-003', 'John', 'Smith', 70, 80, 70, 80, 'pass', 'installment'),
			('1A-004', 'Jane', 'White', 80, 70, 80, 80, 'pass', 'full'),
			('1A-005', 'Abrams', 'White', 80, 70, 80, 80, 'pass', 'full'),
			('1A-006', 'Herdi', 'White', 80, 70, 80, 80, 'fail', 'not paid'),
			('1A-007', 'Wendy', 'White', 100, 95, 80, 80, 'fail', 'installment'),
			('1A-008', 'Ardi', 'White', 100, 95, 80, 80, 'pass', 'not paid'),
			('1A-009', 'Abrams', 'Smith', 95, 93, 80, 80, 'fail', 'not paid'),
			('1A-010', 'Welly', 'White', 95, 93, 80, 80, 'fail', 'not paid'),
			('1B-001', 'Indah', 'Sudarni', 95, 93, 80, 80, 'fail', 'full'),
			('1B-002', 'Aren', 'White', 80, 70, 80, 80, 'pass', 'full'),
			('1B-003', 'John', 'Bernard', 80, 90, 70, 80, 'fail', 'installment'),
			('1B-004', 'Jane', 'Abrams', 90, 80, 90, 80, 'pass', 'full'),
			('1B-005', 'John', 'Albert', 70, 80, 70, 80, 'pass', 'installment');`)

			if err != nil {
				panic("error SQL execute: " + err.Error())
			}

			res, err := main.QuerySQL(dbConn)
			Expect(err).To(BeNil())

			Expect(res).To(Equal([]main.FinalScore{
				{ID: 14, Fullname: "Jane Abrams", Class: "1B", AverageScore: 85},
				{ID: 1, Fullname: "John Doe", Class: "1A", AverageScore: 80},
				{ID: 12, Fullname: "Aren White", Class: "1B", AverageScore: 77},
				{ID: 5, Fullname: "Abrams White", Class: "1A", AverageScore: 77},
				{ID: 4, Fullname: "Jane White", Class: "1A", AverageScore: 77},
			}))
			Expect(len(res)).To(Equal(5))
		})
	})
})
