package rest

import (
	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/zenorachi/literature-list/api/rest/v1"
	"github.com/zenorachi/literature-list/internal/config"
	"github.com/zenorachi/literature-list/internal/service"
	"net/http"
)

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	if ginMode := cfg.GIN.Mode; ginMode != "" {
		gin.SetMode(ginMode)
	}

	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(r *gin.Engine) {
	handlerV1 := v1.NewHandler(h.service)
	api := r.Group("/api")
	{
		handlerV1.Init(api)
	}
}
