package app

import (
	"dashboardapi/modules/auth"
	module_file "dashboardapi/modules/file"
	"dashboardapi/modules/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	auth.AuthRoute(r)
	user.UserRoute(r)
	module_file.Route(r)
	return r
}
