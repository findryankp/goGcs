package goGcs

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func IsExistFile(imagePath string) (*os.File, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return &os.File{}, errors.New("Error opening file " + err.Error())
	}

	return file, nil
}

func LimitMaxSize(file *os.File) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return errors.New("error getting file size " + err.Error())
	}

	fileSize := fileInfo.Size()
	maxFileSize := int64(MaxFileLimit)
	if fileSize > maxFileSize {
		return errors.New("error: file size exceeds maximum limit of 10MB")
	}

	return nil
}

func IsExtentionImage(imagePath string) error {
	fileExtension := filepath.Ext(imagePath)
	fileExtension = strings.ToLower(fileExtension)

	if fileExtension == ".jpg" || fileExtension == ".png" || fileExtension == ".jpeg" {
		return nil
	}

	return errors.New("not image file, only extention jpg,png, and jpeg")
}
