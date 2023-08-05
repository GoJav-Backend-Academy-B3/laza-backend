package helper

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

var DefaultImageUrl = "https://res.cloudinary.com/dmacd1wra/image/upload/v1691249011/1691248119685081467.png"

type CloudinaryCredentials struct {
	Name   string
	APIKey string
	Secret string
}

func GetCloudinaryCredentials() CloudinaryCredentials {
	return CloudinaryCredentials{
		Name:   os.Getenv("CLOUD_NAME"),
		APIKey: os.Getenv("CLOUD_APIKEY"),
		Secret: os.Getenv("CLOUD_SECRET"),
	}
}

func UploadImage(c *gin.Context) (url string, err error) {
	ctx := context.Background()
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 2*1024*1024)

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return DefaultImageUrl, nil
		}
		return
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" && filetype != "image/webp" {
		err = errors.New("file format is not allowed. Please upload a JPEG, JPG, PNG or WEBP image")
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	imageId := strconv.Itoa(int(time.Now().UnixNano()))
	creds := GetCloudinaryCredentials()
	cld, err := cloudinary.NewFromParams(creds.Name, creds.APIKey, creds.Secret)
	if err != nil {
		return
	}

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: imageId})
	if err != nil {
		return
	}

	return result.SecureURL, nil
}
