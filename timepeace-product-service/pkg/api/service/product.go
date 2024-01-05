package service

import (
	"context"
	"fmt"

	"github.com/sgokul961/timepeace-product-service/pkg/helper"
	"github.com/sgokul961/timepeace-product-service/pkg/pb"
	interfaces "github.com/sgokul961/timepeace-product-service/pkg/repository/interface"
)

type ProductService struct {
	Repo interfaces.ProductRepository
	pb.UnimplementedProductServiceServer
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (c *ProductService) AddProduct(ctx context.Context, product *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	sellingPrice := helper.SellingPrice(product.Price, product.Discount)

	err := c.Repo.AddProduct(product, sellingPrice)
	fmt.Println("here", err)
	if err != nil {
		return &pb.AddProductResponse{}, err
	}
	return &pb.AddProductResponse{}, err
}

func (c *ProductService) ListProducts(ctx context.Context, show *pb.ListProductRequest) (*pb.ListProductResponse, error) {

	productResponse, err := c.Repo.ShowAll(show.Page, show.Count)
	if err != nil {
		return nil, err
	}
	updatedProductResponse := make([]*pb.ProductDetail, 0)

	for _, product := range productResponse {
		quantity, _ := c.Repo.Quantity(product.Id)
		if quantity == 0 {
			product.Status = "out of stock"
		} else if quantity == 1 {
			product.Status = "only 1 product remains"
		} else {
			product.Status = "in stock"
		}
		updatedProductResponse = append(updatedProductResponse, product)
	}
	return &pb.ListProductResponse{Products: updatedProductResponse}, nil

}

func (c *ProductService) ProductDetails(ctx context.Context, req *pb.ProductDetailsRequest) (*pb.ProductDetailsResponse, error) {
	response := make([]*pb.FetchProductResponse, 0)
	for _, productId := range req.Id {
		product, err := c.Repo.FetchProduct(productId)
		if err != nil {
			return &pb.ProductDetailsResponse{}, err
		}
		response = append(response, product)
	}
	return &pb.ProductDetailsResponse{Products: response}, nil
}

func (c *ProductService) FetchProduct(ctx context.Context, req *pb.FetchProductRequest) (*pb.FetchProductResponse, error) {
	return c.Repo.FetchProduct(req.Id)
}
