package module_file

import "github.com/gin-gonic/gin"

func Route(app *gin.Engine) {

	r := app.Group("/file")

	r.POST("", UploadImage)
}