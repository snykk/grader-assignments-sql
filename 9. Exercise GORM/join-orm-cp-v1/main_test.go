package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Join with ORM", Ordered, func() {
	dbCredential := main.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
		Schema:       "public",
	}

	dbConn, err := main.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	if err = dbConn.Migrator().DropTable("teachers", "schools", "classes", "lessons"); err != nil {
		panic("failed droping table:" + err.Error())
	}

	BeforeEach(func() {
		dbConn.AutoMigrate(&main.School{}, &main.Class{}, &main.Lesson{}, &main.Teacher{})
		err := main.Reset(dbConn, "schools")
		err = main.Reset(dbConn, "classes")
		err = main.Reset(dbConn, "lessons")
		err = main.Reset(dbConn, "teachers")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Join ORM Test", func() {
		When("schools data initialization", func() {
			It("save the schools data to the postgre database table", func() {
				schoolData := main.School{
					Name:     "SMAN 1 Jakarta",
					Phone:    "(021) 3865001",
					Address:  "Jl. Budi Utomo No.7, Ps. Baru, Kecamatan Sawah Besar, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10710",
					Province: "Jakarta",
				}
				err := schoolData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				result := main.School{}
				dbConn.Model(&main.School{}).First(&result)
				Expect(result.Name).To(Equal(schoolData.Name))
				Expect(result.Phone).To(Equal(schoolData.Phone))
				Expect(result.Address).To(Equal(schoolData.Address))
				Expect(result.Province).To(Equal(schoolData.Province))

				err = main.Reset(dbConn, "schools")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("classes data initialization", func() {
			It("save the classes data to the postgre database table", func() {
				classData := main.Class{Name: "IPA - 1"}
				err := classData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				result := main.Class{}
				dbConn.Model(&main.Class{}).First(&result)
				Expect(result.Name).To(Equal(classData.Name))

				err = main.Reset(dbConn, "classes")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("lessons data initialization", func() {
			It("save the lessons data to the postgre database table", func() {
				lessonData := main.Lesson{Name: "Matematika"}
				err := lessonData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				result := main.Lesson{}
				dbConn.Model(&main.Lesson{}).First(&result)
				Expect(result.Name).To(Equal(lessonData.Name))

				err = main.Reset(dbConn, "lessons")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("teachers data initialization", func() {
			It("save the teachers data to the postgre database table", func() {
				teacherData := main.Teacher{
					Name:     "Aditira",
					Email:    "aditira@gmail.com",
					Phone:    "083831923308",
					SchoolID: 1,
					ClassID:  1,
					LessonID: 1,
				}
				err := teacherData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				result := main.Teacher{}
				dbConn.Model(&main.Teacher{}).First(&result)
				Expect(result.Name).To(Equal(teacherData.Name))

				err = main.Reset(dbConn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("join teacher data with shool, class and lesson data", func() {
			It("display data teachers.name, schools.name, classes.name, lessons.name from each table", func() {
				schoolData := main.School{
					Name:     "SMAN 1 Jakarta",
					Phone:    "(021) 3865001",
					Address:  "Jl. Budi Utomo No.7, Ps. Baru, Kecamatan Sawah Besar, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10710",
					Province: "Jakarta",
				}
				err := schoolData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				classData := main.Class{Name: "IPA - 1"}
				err = classData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				lessonData := main.Lesson{Name: "Matematika"}
				err = lessonData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				teacherData := main.Teacher{
					Name:     "Aditira",
					Email:    "aditira@gmail.com",
					Phone:    "083831923308",
					SchoolID: 1,
					ClassID:  1,
					LessonID: 1,
				}
				err = teacherData.Init(dbConn)
				Expect(err).ShouldNot(HaveOccurred())

				res, err := teacherData.Join(dbConn)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res).To(HaveLen(1))
				Expect(res[0].TeacherName).To(Equal("Aditira"))
				Expect(res[0].SchoolName).To(Equal("SMAN 1 Jakarta"))
				Expect(res[0].ClassName).To(Equal("IPA - 1"))
				Expect(res[0].LessonName).To(Equal("Matematika"))

				err = main.Reset(dbConn, "schools")
				err = main.Reset(dbConn, "classes")
				err = main.Reset(dbConn, "lessons")
				err = main.Reset(dbConn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

	})
})
