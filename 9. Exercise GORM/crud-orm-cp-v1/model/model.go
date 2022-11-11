package model

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique_index"`
	Phone    string
	Address  string
	Province string
}

type Class struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
}

type Lesson struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique_index"`
}

type Teacher struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);unique_index"`
	Email    string
	Phone    string
	LessonID uint
	ClassID  uint
	SchoolID uint
}

type Join struct {
	TeacherName string
	SchoolName  string
	ClassName   string
	LessonName  string
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
