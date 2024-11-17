package auth

import (
	"dashboardapi/config"
	"dashboardapi/db"
	"dashboardapi/db/models"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	db *gorm.DB
}

func AuthService() *authService {
	return &authService{
		db: db.RunDb(),
	}
}

func (s *authService) Login(payload LoginPayload) *LoginResponse {
	user := models.User{
		Username: payload.Username,
	} 
	res := s.db.First(&user)
	if res.RowsAffected == 0 {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err!= nil {
        return nil
    }

	claim := jwt.MapClaims{
		"id": user.ID,
		"name": user.Name,
		"username": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	accessToken, err := token.SignedString(config.Conf.JWT_SECRET)

	if err != nil {
		panic("fail to create JWT token")
	}
	
	return &LoginResponse{
		User: user,
		AccessToken: accessToken,
	}

}

func (s *authService) Register(payload RegisterPayload) (*models.User, error) {

	user := models.User{
		Name: payload.Name,
		Username: payload.Username,
		Password: payload.Password,
	}
	
	res := s.db.Create(&user)

	if res.Error != nil {
		return nil, errors.New("fail to register user")
	}

	return &user, nil;
	
}