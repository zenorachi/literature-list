package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/literature-list/api/rest/utils"
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

// @Summary Search literature list
// @Description Search literature list in cyberleninka
// @Tags literature,search
// @Accept json
// @Produce json
// @Success 200 {object} getTaskByIDResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
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

	utils.SuccessResponse(c, http.StatusOK, res)
}
