package models

import (
	"time"

	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Username  string `gorm:"type:varchar(255);unique_index;not null" json:"username"`
	Password  string `gorm:"type:text;not null;" json:"-"`
	CreatedAt *time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"type:timestamp" json:"updatedAt"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = uuidv7.New().String()
	hashPassword(&u.Password)
	return nil
}

func hashPassword(password *string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	*password = string(hashedPassword)
}
