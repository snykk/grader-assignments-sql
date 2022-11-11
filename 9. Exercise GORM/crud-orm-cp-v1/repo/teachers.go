package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	return t.db.Save(&data).Error
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	err := t.db.Unscoped().Find(&teachers).Error
	return teachers, err // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	return t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name).Error // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	// return t.db.Delete(&model.Teacher{}, int(id)).Error // TODO: replace this
	return t.db.Where("id = ?", id).Delete(&model.Teacher{}).Error
}
