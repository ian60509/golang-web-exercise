package repository
//這裡會實作外層repository.go 定義的interface

import (
	"github.com/jinzhu/gorm"
	"github.com/ian60509/golang-web-exercise/web_basic-5_Intergrate/model"
	"github.com/ian60509/golang-web-exercise/web_basic-5_Intergrate/module/user"
)

type UserRepository struct {
	orm *gorm.DB
}

func NewUserRepository(orm *gorm.DB) user.Repository {
	return &UserRepository{
		orm: orm,
	}
}

func (u *UserRepository) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	var (
		err error
		in  = make([]*model.User, 0)
	)

	err = u.orm.Find(&in, data).Error
	return in, err
}

func (u *UserRepository) GetUser(in *model.User) (*model.User, error) {
	var err error
	err = u.orm.First(&in).Error
	return in, err
}

func (u *UserRepository) CreateUser(in *model.User) (*model.User, error) {
	var err error
	err = u.orm.Create(&in).Error
	return in, err
}

func (u *UserRepository) UpdateUser(in *model.User) (*model.User, error) {
	var err error
	err = u.orm.Save(&in).Error
	return in, err
}

func (u *UserRepository) ModifyUser(in *model.User, data map[string]interface{}) (*model.User, error) {
	var err error
	err = u.orm.Model(&in).Updates(data).Error
	return in, err
}

func (u *UserRepository) DeleteUser(in *model.User) error {
	return u.orm.Delete(&in).Error
}