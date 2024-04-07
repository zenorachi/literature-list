package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/literature-list/internal/service"
)

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initLiteratureRoutes(v1)
	}
}
