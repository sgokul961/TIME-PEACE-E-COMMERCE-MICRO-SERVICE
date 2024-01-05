package interfaces

import "github.com/sgokul961/timepeace-product-service/pkg/pb"

type ProductRepository interface {
	AddProduct(product *pb.AddProductRequest, sellingPrice float64) error
	ShowAll(page, count int32) ([]*pb.ProductDetail, error)
	Quantity(id uint64) (uint32, error)
	FetchProduct(id uint64) (*pb.FetchProductResponse, error)
}
