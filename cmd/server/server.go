package server

import (
	"fmt"
	"net/http"

	"github.com/02amanag/upload-and-download-file/internal/service"
	"github.com/02amanag/upload-and-download-file/internal/usecase"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

var (
	usecaseObject = usecase.NewUsecaseStruct()
	serviceObject = service.NewServiceStruct(usecaseObject)
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := serviceObject.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func Run() {

	server := gin.Default()
	server.Use(CORSMiddleware())

	v1 := server.Group("/group")
	{
		v1.GET("/v1", serviceObject.Healthy)
		v1.GET("/download", serviceObject.DownloadFile)
		v1.POST("/uploadmany", serviceObject.UploadManyFile)
		v1.POST("/uploadone", serviceObject.UploadSingleFile)
		// v1.GET("/download", TokenAuthMiddleware(), serviceObject.DownloadFile)

	}
	server.Run(":8080")
}
