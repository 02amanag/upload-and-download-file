//this is very important sector where all the reusable aur small distinguish function can be wrriten here
package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayError(c *gin.Context, message string, err error) {
	print(message + " <<<- this error @ this endpoint ->>> ") // to print all erroe at console
	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Note": message, "Error": err})
}
