package main_test

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api Todo List", func() {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := db.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	if err = conn.Migrator().DropTable("teachers", "schools", "classes", "lessons"); err != nil {
		panic("failed droping table:" + err.Error())
	}

	schoolRepo := repo.NewSchoolRepo(conn)
	classRepo := repo.NewClassRepo(conn)
	lessonRepo := repo.NewLessonRepo(conn)
	teacherRepo := repo.NewTeacherRepo(conn)

	BeforeEach(func() {
		err := conn.AutoMigrate(&model.School{}, &model.Class{}, &model.Lesson{}, &model.Teacher{})
		err = db.Reset(conn, "teachers")
		err = db.Reset(conn, "schools")
		err = db.Reset(conn, "classes")
		err = db.Reset(conn, "lessons")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Repository", func() {
		Describe("Schools", func() {
			When("schools data initialization", func() {
				It("save the schools data to the postgre database table in the form of a list", func() {
					schoolData := []model.School{
						{
							Name:     "SMAN 1 Jakarta",
							Phone:    "(021) 3865001",
							Address:  "Jl. Budi Utomo No.7, Ps. Baru, Kecamatan Sawah Besar, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10710",
							Province: "Jakarta",
						},
						{
							Name:     "SMA Negeri 1 Depok",
							Phone:    "(021) 7520137",
							Address:  "Jl. Nusantara Raya No.317, Depok Jaya, Kec. Pancoran Mas, Kota Depok, Jawa Barat 16432",
							Province: "Jawa Barat",
						},
					}
					err := schoolRepo.Init(schoolData)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.School{}
					conn.First(&model.School{}).First(&result)
					Expect(result.Name).To(Equal(schoolData[0].Name))
					Expect(result.Phone).To(Equal(schoolData[0].Phone))
					Expect(result.Address).To(Equal(schoolData[0].Address))
					Expect(result.Province).To(Equal(schoolData[0].Province))

					err = db.Reset(conn, "schools")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Classes", func() {
			When("classes data initialization", func() {
				It("save the classes data to the postgre database table in the form of a list", func() {
					classData := []model.Class{{Name: "IPA - 1"}, {Name: "IPA - 2"}, {Name: "IPA - 3"}, {Name: "IPS - 1"}, {Name: "IPS - 2"}, {Name: "IPS - 3"}}
					err := classRepo.Init(classData)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Class{}
					conn.First(&model.Class{}).First(&result)
					Expect(result.Name).To(Equal(classData[0].Name))

					err = db.Reset(conn, "classes")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Lessons", func() {
			When("lessons data initialization", func() {
				It("save the lessons data to the postgre database table in the form of a list", func() {
					lessonData := []model.Lesson{{Name: "Matematika"}, {Name: "Biologi"}, {Name: "Fisika"}, {Name: "Kimia"}, {Name: "Geografi"}, {Name: "Ekonomi"}, {Name: "Sejarah"}}
					err := lessonRepo.Init(lessonData)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Lesson{}
					conn.First(&model.Lesson{}).First(&result)
					Expect(result.Name).To(Equal(lessonData[0].Name))

					err = db.Reset(conn, "lessons")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Teachers", func() {
			When("save teacher data", func() {
				It("save teacher data to postgre database with teachers table", func() {
					teacher := model.Teacher{Name: "Aditira", Email: "aditira@gmail.com", Phone: "08334232322", SchoolID: uint(1), ClassID: uint(1), LessonID: uint(3)}
					err := teacherRepo.Save(teacher)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Teacher{}
					conn.First(&model.Teacher{}).First(&result)
					Expect(result.Name).To(Equal(teacher.Name))
					Expect(result.Email).To(Equal(teacher.Email))
					Expect(result.Phone).To(Equal(teacher.Phone))
					Expect(result.SchoolID).To(Equal(teacher.SchoolID))
					Expect(result.ClassID).To(Equal(teacher.ClassID))
					Expect(result.LessonID).To(Equal(teacher.LessonID))

					err = db.Reset(conn, "teachers")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("query teacher data", func() {
				It("display teacher data to postgre database with teachers table", func() {
					teacher := model.Teacher{Name: "Aditira", Email: "aditira@gmail.com", Phone: "08334232322", SchoolID: uint(1), ClassID: uint(1), LessonID: uint(3)}
					err := teacherRepo.Save(teacher)
					Expect(err).ShouldNot(HaveOccurred())

					res, err := teacherRepo.Query()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res).To(HaveLen(1))
					Expect(res[0].Name).To(Equal(teacher.Name))
					Expect(res[0].Email).To(Equal(teacher.Email))
					Expect(res[0].Phone).To(Equal(teacher.Phone))
					Expect(res[0].SchoolID).To(Equal(teacher.SchoolID))
					Expect(res[0].ClassID).To(Equal(teacher.ClassID))
					Expect(res[0].LessonID).To(Equal(teacher.LessonID))
					Expect(res[0].Model.DeletedAt.Valid).To(BeFalse())

					err = db.Reset(conn, "teachers")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("update teacher data", func() {
				It("update teacher data to postgre database with teachers table", func() {
					teacher := model.Teacher{Name: "Aditira", Email: "aditira@gmail.com", Phone: "08334232322", SchoolID: uint(1), ClassID: uint(1), LessonID: uint(3)}
					err := teacherRepo.Save(teacher)
					Expect(err).ShouldNot(HaveOccurred())

					err = teacherRepo.Update(uint(1), "Aditira Jamhuri")
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Teacher{}
					conn.First(&model.Teacher{}).First(&result)
					Expect(result.Name).To(Equal("Aditira Jamhuri"))

					err = db.Reset(conn, "teachers")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("delete teacher data", func() {
				It("delete teacher data to postgre database with teachers table", func() {
					teacher := model.Teacher{Name: "Aditira", Email: "aditira@gmail.com", Phone: "08334232322", SchoolID: uint(1), ClassID: uint(1), LessonID: uint(3)}
					err := teacherRepo.Save(teacher)
					Expect(err).ShouldNot(HaveOccurred())

					err = teacherRepo.Delete(uint(1))
					Expect(err).ShouldNot(HaveOccurred())

					res, err := teacherRepo.Query()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res[0].Model.DeletedAt.Valid).To(BeTrue())

					err = db.Reset(conn, "teachers")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
