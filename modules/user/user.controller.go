package user

import (
	"dashboardapi/utils"

	"github.com/gin-gonic/gin"
)

func GetAllusers(c *gin.Context) {

	s := UserService()

	user := s.GetAllUser()
	utils.ResOkFormater(c, "success", user, 200)
}
