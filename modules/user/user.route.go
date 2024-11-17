package user

import "github.com/gin-gonic/gin"

func UserRoute(r *gin.Engine) {
	user := r.Group("/users")
	user.GET("/", GetAllusers)
}