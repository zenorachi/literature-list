package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/literature-list/api/rest/utils"
	"github.com/zenorachi/literature-list/internal/models"
	"net/http"
)

func (h *Handler) initLiteratureRoutes(api *gin.RouterGroup) {
	literature := api.Group("/literature")
	{
		literature.GET("/search", h.searchLiterature)
	}
}

type searchLiteratureRequest struct {
	LiteratureList []string `json:"literature_list" binding:"required"`
}

type searchLiteratureResponse struct {
	LiteratureList []models.Literature `json:"literature_list"`
}

// @Summary Search literature list
// @Description Search literature list in cyberleninka
// @Tags literature,search
// @Accept json
// @Produce json
// @Success 200 {object} searchLiteratureResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/literature/search [get]
func (h *Handler) searchLiterature(c *gin.Context) {
	var req searchLiteratureRequest
	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.service.SearchLiterature(req.LiteratureList)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, searchLiteratureResponse{
		LiteratureList: res,
	})
}
