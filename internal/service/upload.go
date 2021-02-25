package service

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/02amanag/upload-and-download-file/internal/helper"

	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) UploadManyFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	path, err := filepath.Abs("file/save")
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading Multiple file , keep the key = 'file' to avoid this error ", err)
	}
	for _, file := range files {
		log.Println(file.Filename)
		err := ctx.SaveUploadedFile(file, path+"/"+file.Filename)
		if err != nil {
			helper.DisplayError(ctx, "Error in Uploading Multiple file ,as the storage dir. not found ", err)
		}
	}

	fmt.Println(ctx.PostForm("key"))
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func (s *ServiceStruct) UploadSingleFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading file, keep the key = 'file' to avoid this error ", err)
	}
	log.Println(file.Filename)

	path, err := filepath.Abs("file/save")
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading file, as the storage dir. not found ", err)
	}

	err = ctx.SaveUploadedFile(file, path+"/"+file.Filename)
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading Multiple file, as the storage dir. not found ", err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
