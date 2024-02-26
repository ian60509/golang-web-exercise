package user

import "github.com/ian60509/golang-web-exercise/web_basic-5_Intergrate/model"

type Repository interface {
	GetUserList() (map[string]interface{}) ([]*model.User, error)
	GetUser(in *model.User) (*model.User, error)
	CreateUser(in *model.User) (*model.User, error)
	ModifyUser(in *model.User, data map[string]interface{}) (*model.User, error)
	DeleteUser(in *model.User) error
}