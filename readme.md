# Simple Google Cloud Storage - Golang
Simple package to manage your image in Google Cloud Storage

**Development by:** 
- Findryankp

## Import
```shell
go get github.com/Findryankp/goGcs
```

## Example
```go
func main() {
	// set configuration
	MaxFileLimit = 10000000
	GscCredentialFilename = "filecred.json"

	bucketName := "bucket_images"
	filePathName := "testImage.png"
	fileNameInGCS := "testImageBucket.png"

    //upload file
	if err := UploadImage(bucketName, filePathName, fileNameInGCS); err != nil {
		fmt.Println(err.Error())
	}

    //delete file
	if err := DeleteImage(bucketName, filePathName); err != nil {
		fmt.Println(err.Error())
	}
}
```

## Tutorial
1. Create Your Bucket First
2. Generate Key