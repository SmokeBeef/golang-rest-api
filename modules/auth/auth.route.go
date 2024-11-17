package auth

import "github.com/gin-gonic/gin"

func AuthRoute(r *gin.Engine) {
	gr := r.Group("/auth")
	gr.POST("/login", Login)
	gr.POST("/register", Register)
}