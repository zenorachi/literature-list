package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/literature-list/pkg/logger"
)

func ErrorResponse(c *gin.Context, code int, err error) {
	logger.Error(err.Error(), fmt.Sprintf("code: %d", code))
	c.AbortWithStatusJSON(code, err.Error())
}

func SuccessResponse(c *gin.Context, code int, obj any) {
	logger.Info(fmt.Sprintf("code: %d", code))
	c.JSON(code, obj)
}
