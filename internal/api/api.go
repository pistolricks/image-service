package api

import "mime/multipart"

type RequestBody struct {
	File     multipart.File `validate:"required"`
	FileName string         `validate:"required"`
	Format   string         `validate:"required,oneof=avif webp"`
	Quality  int
	Width    int
	Height   int
}

type ResponseBody struct {
	File        []byte
	FileName    string
	ContentType string
}
