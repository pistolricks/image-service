package api

import (
	"github.com/h2non/bimg"
	"io"
	"strings"
)

func ConvertImage(payload *RequestBody) (*ResponseBody, error) {
	fileContent, _ := io.ReadAll(payload.File)
	fileType := getFileType(payload.Format)
	fileName := getFileName(payload.FileName, payload.Format)
	options := bimg.Options{
		Quality: payload.Quality,
		Width:   payload.Width,
		Height:  payload.Height,
	}

	image, err := bimg.NewImage(fileContent).Process(options)
	if err != nil {
		return nil, err
	}

	image, err = bimg.NewImage(image).Convert(fileType)
	if err != nil {
		return nil, err
	}

	return &ResponseBody{
		File:        image,
		FileName:    fileName,
		ContentType: payload.Format,
	}, nil
}

func getFileType(format string) bimg.ImageType {
	if format == "avif" {
		return bimg.AVIF
	} else {
		return bimg.WEBP
	}
}

func getFileName(originalFileName string, format string) string {
	splits := strings.Split(originalFileName, ".")
	splits[len(splits)-1] = format
	return strings.Join(splits, ".")
}
