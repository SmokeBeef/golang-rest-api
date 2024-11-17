package module_file

import (
	lib_minio "dashboardapi/lib/minio"
	"mime/multipart"
)

type fileService struct {
}

func Service() *fileService {

	return &fileService{}
}

func (s *fileService) UploadFile(file *multipart.FileHeader) (string, error) {

	fileName := file.Filename
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileContent.Close()

	contentType := file.Header["Content-Type"][0]

	l, err := lib_minio.Run().UploadFile(fileName, fileContent, file.Size, contentType)

	if err != nil {
		return "", err
	}

	return l, nil

}
