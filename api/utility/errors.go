package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinCtxError(c *gin.Context, errMsg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": errMsg,
	})

}
