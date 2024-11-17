package auth

import "dashboardapi/db/models"

type LoginPayload struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string      `json:"accessToken"`
	User        models.User `json:"user"`
}

type RegisterPayload struct {
	Name     string `json:"name" binding:"required" form:"name"`
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}
