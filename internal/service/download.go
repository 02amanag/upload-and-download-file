package service

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) DownloadFile(ctx *gin.Context) {
	path, err := filepath.Abs("file/give")
	if err != nil {
		log.Fatal(err)
	}
	fileName := "sample.pdf"

	ctx.FileAttachment(path+"/"+fileName, fileName)
}
