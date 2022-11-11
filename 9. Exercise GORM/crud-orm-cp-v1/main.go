package main

import (
	"fmt"

	_ "embed"

	_ "github.com/lib/pq"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repo"
	"a21hc3NpZ25tZW50/terminal"
)

func main() {
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
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.School{}, &model.Class{}, &model.Lesson{}, &model.Teacher{})

	schoolRepo := repo.NewSchoolRepo(conn)
	classRepo := repo.NewClassRepo(conn)
	lessonRepo := repo.NewLessonRepo(conn)
	teacherRepo := repo.NewTeacherRepo(conn)
	teacherTerminal := terminal.NewTeacherTerminal(teacherRepo, conn)

	school := []model.School{
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
	schoolRepo.Init(school)
	class := []model.Class{{Name: "IPA - 1"}, {Name: "IPA - 2"}, {Name: "IPA - 3"}, {Name: "IPS - 1"}, {Name: "IPS - 2"}, {Name: "IPS - 3"}}
	classRepo.Init(class)
	lesson := []model.Lesson{{Name: "Matematika"}, {Name: "Biologi"}, {Name: "Fisika"}, {Name: "Kimia"}, {Name: "Geografi"}, {Name: "Ekonomi"}, {Name: "Sejarah"}}
	lessonRepo.Init(lesson)

	for {
		fmt.Println("Select Action:")
		fmt.Println("1. Go to Add Teacher App")
		fmt.Println("Or submit any key to Logout!")
		fmt.Printf("Please your action: ")
		var choice int
		fmt.Scan(&choice)
		fmt.Printf("\x1bc")

		if choice == 1 {
			teacherTerminal.TeacherApp()
		} else {
			fmt.Println("Logged Out")
			break
		}
	}

	db.Reset(conn, "schools")
	db.Reset(conn, "classes")
	db.Reset(conn, "lessons")
}
