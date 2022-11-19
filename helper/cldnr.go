package helper

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUpload(image interface{}, folder string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(LoadEnvFile("CLOUDINARY_CLOUD_NAME"), LoadEnvFile("CLOUDINARY_API_KEY"), LoadEnvFile("CLOUDINARY_API_SECRECT_KEY"))
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, image, uploader.UploadParams{Folder: folder})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

func DeleteImage(pubId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(LoadEnvFile("CLOUDINARY_CLOUD_NAME"), LoadEnvFile("CLOUDINARY_API_KEY"), LoadEnvFile("CLOUDINARY_API_SECRECT_KEY"))
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: pubId})
	if err != nil {
		return "", err
	}

	return resp.Result, nil
}
