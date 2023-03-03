package rest

import "github.com/gin-gonic/gin"

func (h *rest) GetAllCategories(c *gin.Context) {
	categories,status,err:=h.uc.Category.FindAll()
	if err!=nil{
		h.ErrorResponse(c,status,err,"Failed to querying categories data")
		return
	}
	
	h.SuccessResponse(c,status,"Querying speaker's categories data successfully",categories)
}