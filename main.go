package response

import (
	"github.com/gin-gonic/gin"
)

type R struct {
	responseCode int
	errCode      string
	message      string
}

func BadRequest(c *gin.Context, errCode string) {
	var r R
	r.errCode = errCode
	r.message = "Request failed"
	r.responseCode = 400
	c.AbortWithStatusJSON(r.responseCode, gin.H{"message": r.message, "errCode": r.errCode})
}

func Unauthorized(c *gin.Context, message string) {
	var r R
	r.errCode = "ERR419"
	r.message = "Unauthorized"
	r.responseCode = 419
	c.AbortWithStatusJSON(r.responseCode, gin.H{"message": r.message, "errCode": r.errCode, "errMessage": message})
}

func InternalError(c *gin.Context, message string) {
	var r R
	r.errCode = "ERR500"
	r.message = "Something went wrong"
	r.responseCode = 500
	c.AbortWithStatusJSON(r.responseCode, gin.H{"message": r.message, "errCode": r.errCode, "errMessage": message})
}

func AccessDenied(c *gin.Context, message string) {
	var r R
	r.errCode = "ERR405"
	r.message = "Access denied"
	r.responseCode = 500
	c.AbortWithStatusJSON(r.responseCode, gin.H{"message": r.message, "errCode": r.errCode, "errMessage": message})
}

func ToJson(c *gin.Context, h gin.H) {
	var r R
	r.responseCode = 200
	c.JSON(200, h)
}
