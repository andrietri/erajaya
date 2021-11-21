package product

import (
	"erajaya/entities/model"
	"erajaya/usecase/product"
	"erajaya/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	productUsecase product.ProductUsecase
}

func NewProductHandler(productUsecase product.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) GetProductList(c *gin.Context) {
	ctx := c.Request.Context()
	sortBy, _ := c.GetQuery("sortBy")
	sortType, _ := c.GetQuery("sortType")

	res, err := h.productUsecase.Fetch(ctx, sortBy, sortType)
	if err != nil {
		if err.Error() == "data not found" {
			utils.Response(c, 404, err.Error(), nil)
			return
		}
		utils.Response(c, 500, "something went wrong", nil)
		return
	}

	utils.Response(c, 200, "success", res)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	validate = validator.New()
	product := model.ProductRequest{}

	err := c.ShouldBindJSON(&product)
	if err != nil {
		utils.Response(c, 400, err.Error(), nil)
		return
	}

	err = validate.Struct(product)
	if err != nil {
		utils.Response(c, 400, err.Error(), nil)
		return
	}

	ctx := c.Request.Context()
	res, err := h.productUsecase.Store(ctx, product)
	if err != nil {
		utils.Response(c, 500, "something went wrong", nil)
		return
	}

	utils.Response(c, 201, "success", res)
}
