package product

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/product"
	"github.com/williamchang80/sea-apd/dto/domain"
	request "github.com/williamchang80/sea-apd/dto/request/product"
	"github.com/williamchang80/sea-apd/dto/response/base"
	response "github.com/williamchang80/sea-apd/dto/response/product"
	"net/http"
)

type ProductController struct {
	usecase product.ProductUsecase
}

func NewProductController(e *echo.Echo, p product.ProductUsecase) product.ProductController {
	c := &ProductController{
		usecase: p,
	}
	e.GET("/api/products", c.GetProducts)
	e.POST("/api/product", c.CreateProduct)
	e.GET("/api/product", c.GetProductById)
	e.PUT("/api/product", c.UpdateProduct)
	e.DELETE("/api/product", c.DeleteProduct)
	e.GET("/api/merchant/products", c.GetProductsByMerchant)
	return c
}

func (p *ProductController) GetProducts(c echo.Context) error {
	products, err := p.usecase.GetProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}

	return c.JSON(http.StatusOK, &response.GetProductsResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.ProductListDto{Products: products},
	})
}

func (p *ProductController) CreateProduct(c echo.Context) error {
	var productRequest request.ProductRequest
	c.Bind(&productRequest)
	if err := p.usecase.CreateProduct(productRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: message.BAD_REQUEST,
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusCreated,
		Message: message.SUCCESS,
	})
}

func (p *ProductController) GetProductById(context echo.Context) error {
	id := context.QueryParam("product_id")
	product, err := p.usecase.GetProductById(id)
	if err != nil {
		return context.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return context.JSON(http.StatusOK, &response.GetProductByIdResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.ProductDto{
			Product: product,
		},
	})
}

func (p *ProductController) UpdateProduct(context echo.Context) error {
	var productRequest request.ProductRequest
	context.Bind(&productRequest)
	productId := context.FormValue("productId")
	err := p.usecase.UpdateProduct(productId, productRequest)
	if err != nil {
		return context.JSON(http.StatusBadRequest, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: message.BAD_REQUEST,
		})
	}
	return context.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (p *ProductController) DeleteProduct(context echo.Context) error {
	id := context.QueryParam("productId")
	err := p.usecase.DeleteProduct(id)
	if err != nil {
		return context.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return context.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (p *ProductController) GetProductsByMerchant(c echo.Context) error {
	merchantId := c.QueryParam("merchantId")
	products, err := p.usecase.GetProductsByMerchant(merchantId)
	if err != nil {
		c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}

	return c.JSON(http.StatusOK, &response.GetProductsResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.ProductListDto{Products: products},
	})
}
