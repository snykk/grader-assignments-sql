package terminal

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repo"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type TeacherTerminal struct {
	teacherRepo repo.TeacherRepo
	Db          *gorm.DB
}

func NewTeacherTerminal(teacherRepo repo.TeacherRepo, db *gorm.DB) TeacherTerminal {
	return TeacherTerminal{teacherRepo, db}
}

func (t TeacherTerminal) TeacherApp() error {
	var err error
	for {
		teachers, err := t.teacherRepo.Query()
		if err != nil {
			return err
		}

		fmt.Println("Query Teacher: ")
		for i, v := range teachers {
			fmt.Println(i+1, "Name: ", v.Name, "Email: ", v.Email, "Phone: ", v.Phone, "SchoolID: ", v.SchoolID, "ClassID: ", v.ClassID, "LessonID: ", v.LessonID, "Updated At: ", v.UpdatedAt, "Deleted At: ", v.DeletedAt)
		}

		fmt.Println("Welcome to Teacher App")
		fmt.Println("Select menu:")
		fmt.Println("1. Add Teacher")
		fmt.Println("2. Update Teacher")
		fmt.Println("3. Delete Teacher")
		fmt.Println("submit any key to exit!")
		fmt.Printf("Please submit menu: ")
		var choice int
		fmt.Scan(&choice)
		fmt.Printf("\x1bc")

		if choice == 1 {
			err = t.TeacherAdd()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Teacher Added")
			}
		} else if choice == 2 {
			err = t.TeacherUpdate()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Teacher Name Updated")
			}
		} else if choice == 3 {
			err = t.TeacherDelete()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Teacher Deleted")
			}
		} else {
			fmt.Println("Invalid menu!")
		}

		var again string
		fmt.Printf("Add teacher again? (y/n): ")
		fmt.Scan(&again)

		if again == "y" {
			fmt.Printf("\x1bc")
			continue
		} else {
			fmt.Printf("\x1bc")
			break
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func (t TeacherTerminal) TeacherAdd() error {
	teacher := model.Teacher{}

	q := []string{"Name: ", "Email: ", "Phone: ", "SchoolID: ", "ClassID: ", "LessonID: "}
	res := []string{}
	for _, v := range q {
		fmt.Printf(v)
		var a string
		fmt.Scan(&a)
		res = append(res, a)
	}

	schoolID, _ := strconv.Atoi(res[3])
	classID, _ := strconv.Atoi(res[4])
	lessonID, _ := strconv.Atoi(res[5])
	teacher = model.Teacher{Name: res[0], Email: res[1], Phone: res[2], SchoolID: uint(schoolID), ClassID: uint(classID), LessonID: uint(lessonID)}
	err := t.teacherRepo.Save(teacher)

	if err != nil {
		return err
	}

	return nil
}

func (t TeacherTerminal) TeacherUpdate() error {
	teachers, err := t.teacherRepo.Query()
	if err != nil {
		return err
	}

	fmt.Println("Query Teacher: ")
	for i, v := range teachers {
		fmt.Println(i+1, "Name: ", v.Name, "Email: ", v.Email, "Phone: ", v.Phone, "SchoolID: ", v.SchoolID, "ClassID: ", v.ClassID, "LessonID: ", v.LessonID)
	}

	fmt.Printf("select id to update name: ")
	var id string
	fmt.Scan(&id)
	idInt, _ := strconv.Atoi(id)
	fmt.Printf("write new name: ")
	var name string
	fmt.Scan(&name)

	err = t.teacherRepo.Update(uint(idInt), name)
	if err != nil {
		return err
	}
	return nil
}

func (t TeacherTerminal) TeacherDelete() error {
	teachers, err := t.teacherRepo.Query()
	if err != nil {
		return err
	}

	fmt.Println("halo")
	fmt.Println("Query Teacher: ")
	for i, v := range teachers {
		fmt.Println(i+1, "Name: ", v.Name, "Email: ", v.Email, "Phone: ", v.Phone, "SchoolID: ", v.SchoolID, "ClassID: ", v.ClassID, "LessonID: ", v.LessonID, "Deleted At: ", v.DeletedAt)
	}

	fmt.Printf("select id to delete record: ")
	var id string
	fmt.Scan(&id)
	idInt, _ := strconv.Atoi(id)

	err = t.teacherRepo.Delete(uint(idInt))
	if err != nil {
		return err
	}
	return nil
}
