package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/literature-list/pkg/logger"
)

type Error struct {
	Detail string `json:"detail"`
}

func ErrorResponse(c *gin.Context, code int, err error) {
	logger.Error(err.Error(), fmt.Sprintf("code: %d", code))
	c.AbortWithStatusJSON(code, Error{Detail: err.Error()})
}

func SuccessResponse(c *gin.Context, code int, obj any) {
	logger.Info(fmt.Sprintf("code: %d", code))
	c.JSON(code, obj)
}
