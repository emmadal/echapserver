package helpers

import (
	"context"
	"mime/multipart"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// UploadHelper upload file to cloudinary cloud
func UploadHelper(file *multipart.FileHeader) (string, error) {
	cloudName := EnvCloudName()
	key := EnvCloudAPIKey()
	secret := EnvCloudSecretAPI()
	cld, err := cloudinary.NewFromParams(cloudName, key, secret)
	cld.Config.URL.Secure = true
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	// upload file to cloudinary
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}
	url := result.SecureURL // get the secure https image url 
	return url, err
}
