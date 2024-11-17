package user

import (
	"dashboardapi/db"
	"dashboardapi/db/models"

	"gorm.io/gorm"
)

type userService struct {
	DB *gorm.DB
}

func UserService() *userService {
	return &userService{
		DB: db.DB,
	}
}

func (s *userService) GetAllUser() []models.User {
	var user []models.User
	s.DB.Limit(10).Find(&user)
	
	return user
}
