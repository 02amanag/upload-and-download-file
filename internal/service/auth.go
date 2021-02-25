package service

import (
	"github.com/02amanag/upload-and-download-file/internal/helper"
	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) TokenValid(c *gin.Context) error {

	_, err := s.usecase.ExtractTokenMetadata(c.Request)
	if err != nil {
		//Token either expired or not valid
		helper.DisplayError(c, "401", err)
		return nil
	}
	return err

}
