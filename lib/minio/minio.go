package lib_minio

import (
	"context"
	"dashboardapi/config"
	"fmt"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioLib struct {
	client *minio.Client
}

var bucketName string = "deva"

func Run() *MinioLib {
	useSSL := true

	// Initialize minio client object.
	log.Printf(config.Conf.S3_ACCESS_KEY, config.Conf.S3_SECRET_KEY)
	minioClient, err := minio.New(config.Conf.S3_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.S3_ACCESS_KEY, config.Conf.S3_SECRET_KEY, ""),
		Secure: useSSL,
		Region: "id-jkt-1-id",
		BucketLookup: minio.BucketLookupPath,
	})
	if err != nil {
		log.Fatalln(err)
	}

	BucketExist, errBucket := minioClient.BucketExists(context.Background(), bucketName)

	if errBucket != nil {
		log.Fatalln(errBucket)
	}
	if !BucketExist {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Successfully created %s bucket\n", bucketName)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
	return &MinioLib{
		client: minioClient,
	}
}

func (mc *MinioLib) UploadFile(name string, file io.Reader, fileSize int64, contentType string) (string, error) {
	log.Printf("Starting to upload %s... \n", name)
	info, err := mc.client.PutObject(context.Background(), bucketName, name, file, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		log.Fatalf("Error uploading file %s to bucket %s: %v\n", name, bucketName, err)
		return "", err
	}

	u, err := mc.client.PresignedGetObject(context.Background(), bucketName, name, time.Hour*5, url.Values{})

	if err != nil {
		log.Fatalf("Error generating presigned URL for object %s in bucket %s: %v\n", name, bucketName, err)
        return "", err
	}

	fmt.Printf("Url ", u.String())
	fmt.Printf("info", info)
	return u.String(), nil
}


