package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/timepeace-api-gateway/pkg/client"
	"github.com/sgokul961/timepeace-api-gateway/pkg/pb"
	"github.com/sgokul961/timepeace-api-gateway/pkg/utils/response"
)

type ProductHandler struct {
	client *client.ProductClient
}

func NewProductHandler(client *client.ProductClient) *ProductHandler {
	return &ProductHandler{client: client}
}

func (p *ProductHandler) AddProduct(c *gin.Context) {
	var product *pb.AddProductRequest

	if err := c.ShouldBindJSON(&product); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	_, err := p.client.Client.AddProduct(context.Background(), product)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully added product", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

func (p *ProductHandler) ListProducts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	count, err := strconv.Atoi(c.DefaultQuery("count", "4"))
	if err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	productDetails, err := p.client.Client.ListProducts(context.Background(), &pb.ListProductRequest{Page: int32(page), Count: int32(count)})

	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing products", productDetails, nil)
	c.JSON(http.StatusOK, succRes)
}
