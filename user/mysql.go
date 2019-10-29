package user

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	DB *gorm.DB
}
func NewMysqlUserRepository(db *gorm.DB) UserRepository {

	// Migrate Models
	db.AutoMigrate(&User{})

	return &userRepository {
	db,
	}
}

func (r *userRepository) Find(id string) (*User, error) {
	user := new(User)

	if r.DB.Where("id = ?", id).First(&user).RecordNotFound() {
		return nil, errors.New("No User Found")
	} else {
		return user, nil
	}
}

func (r *userRepository) FindAll() (users []*User, err error) {

	users = append(users, nil)

	r.DB.Find(&users)

	return users, nil

}

func (r *userRepository) Create(user User) error {

	err := r.DB.Create(&user).Error
	if err != nil {
		return errors.New("Internal Server Error")
	}
	return nil
}

func (r *userRepository) Delete(id string) error {

	var user User

	if r.DB.Where("id = ?", id).First(&user).RecordNotFound() {

		return errors.New("No User Found")

	} else {
		r.DB.Delete(&user, id)
	}

	return nil

}