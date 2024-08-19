package repository

import (
	"gorm.io/gorm"
	"user_server/model"
)

type UserRepository interface {
	Create (user *model.User) error
	GetAll () ([]model.User, error)
	GetById (id int) (model.User, error)
	Update (user *model.User) error
	Delete (user *model.User) error
}

// --------------Implement 物件-------------------------
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB // 如此可以讓 userRepository method 使用 DB 了
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error //如果沒有指定條件，則會取得所有的資料
	return users, err
}

func (r *userRepository) GetById(id int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *model.User) error {
	return r.db.Delete(user).Error
}

