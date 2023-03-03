package rest

import (
	"net/http"
	"project-intern-bcc/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *rest) GetAllSponsors(c *gin.Context) {
	category := c.Query("category")
	keyword := c.Query("keyword")
	var pagination entity.Pagination
	err := c.ShouldBindWith(&pagination, binding.Query)
	if err != nil {
		h.ErrorResponse(c, http.StatusBadRequest, err, "Failed to read parameters")
		return
	}

	result, statusCode, err := h.uc.Company.FindAll(category, keyword, pagination)
	if err != nil {
		h.ErrorResponse(c, statusCode, err, result)
		return
	}
	h.SuccessResponse(c, statusCode, "Querying sponsors data successfully", result)
}