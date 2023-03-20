package goGcs

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func DeleteImage(bucketName, filePathAndName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(GscCredentialFilename))
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	image := bucket.Object(filePathAndName)

	if err := image.Delete(ctx); err != nil {
		return err
	}

	return nil
}
