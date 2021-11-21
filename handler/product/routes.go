package product

import (
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) Route(app *gin.Engine) {
	v1 := app.Group("api/v1")
	v1.Use()
	{
		v1.GET("/products", h.GetProductList)
		v1.POST("/products", h.CreateProduct)
	}
}
