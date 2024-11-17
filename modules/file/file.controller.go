package module_file

import (
	"dashboardapi/utils"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {

	file, err := c.FormFile("file")
	
	if err != nil {
		utils.ResFailFormater(c, "file is required", err.Error(), 400)
		return
	}

	if file == nil {
		utils.ResFailFormater(c, "file is required", nil, 400)
		return
	}

	s, err := Service().UploadFile(file)

	if err!= nil {
        utils.ResFailFormater(c, "error uploading file", err, 500)
        return
    }

	utils.ResOkFormater(c, "file uploaded successfully", s, 200)
}
