package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/tobi007/angular-go-serve/models"
	"github.com/tobi007/angular-go-serve/repository"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewuserDao(Conn *gorm.DB) repository.UserRepo {
	return &userDao{
		Conn: Conn,
	}
}

type userDao struct {
	Conn *gorm.DB
}

func (userRepo *userDao) Create( tpu *models.User) (int64, error) {
	res := userRepo.Conn.Save(tpu)
	return res.RowsAffected, res.Error
}

func (userRepo *userDao) Update(u *models.User) (*models.User, error) {
	return new(models.User), nil
}

func (userRepo *userDao) RetrieveByID (id int64) (*models.User, error)  {
	var user models.User
	if err := userRepo.Conn.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *userDao) RetrieveByEmail (email string) (*models.User, error)  {
	var user models.User
	if err := userRepo.Conn.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *userDao) ExistByID(email string) bool  {
	var user models.User
	if err := userRepo.Conn.Where("email=?", email).First(&user).Error; err != nil && &user == nil {
		return false
	}

	return true
}

func (userRepo *userDao)ConnectionActive() error  {
	return userRepo.Conn.DB().Ping()
}
