package auth

import (
	"dashboardapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var payload LoginPayload
	
	if err := c.ShouldBind(&payload); err != nil {
		err := utils.GetErrValidation(err)
		utils.ResFailFormater(c, "payload not complete", err, http.StatusBadRequest)
		return
	}

	s := AuthService()
	res := s.Login(payload)

	if res == nil {
		utils.ResFailFormater(c, "username or password is incorrect", res, http.StatusUnauthorized)
		return
	}

	utils.ResOkFormater(c, "success login", res, http.StatusOK)
}


func Register(c *gin.Context) {
	var payload RegisterPayload
	err := c.ShouldBind(&payload)
	if err != nil {
		err := utils.GetErrValidation(err)
		utils.ResFailFormater(c, "payload not complete", err, http.StatusBadRequest)
		return
	}
    s := AuthService()
    res, err := s.Register(payload)

	if err != nil {
		utils.ResFailFormater(c, "failed to register", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResOkFormater(c, "success Register", res, http.StatusCreated)
}