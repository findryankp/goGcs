package goGcs

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadImage(bucketName, filePathAndName, gcsFilename string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(GscCredentialFilename))
	if err != nil {
		return err
	}
	defer client.Close()

	file, err := IsExistFile(filePathAndName)
	if err != nil {
		return err
	}
	defer file.Close()

	//check limit
	if err1 := LimitMaxSize(file); err1 != nil {
		return err1
	}

	//check ext
	if err1 := IsExtentionImage(filePathAndName); err1 != nil {
		return err1
	}

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(gcsFilename)
	writer := obj.NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	return nil
}
