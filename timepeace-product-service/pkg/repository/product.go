package repository

import (
	"github.com/sgokul961/timepeace-product-service/pkg/pb"
	interfaces "github.com/sgokul961/timepeace-product-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &ProductRepository{db}
}

func (c *ProductRepository) AddProduct(product *pb.AddProductRequest, sellingPrice float64) error {
	var id uint
	err := c.DB.Raw(`insert into products (name,description,quantity,price,selling_price,discount,category_id,brand_id) values(?,?,?,?,?,?,?,?) RETURNING id`, product.Name, product.Description, product.Quantity, product.Price, sellingPrice, product.Discount, product.Category, product.Brand).Scan(&id).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductRepository) ShowAll(page, count int32) ([]*pb.ProductDetail, error) {
	offset := (page - 1) * count

	var productResponse []*pb.ProductDetail
	if err := c.DB.Raw(`
        SELECT
            p.id,
            p.name,
            p.description,
            p.price,
            p.selling_price,
            
            p.discount,
            c.name AS category,
            b.name AS brand
        FROM
            products p
        JOIN
            categories c ON c.id = p.category_id
        JOIN
            brands b ON b.id = p.brand_id
        
        
        LIMIT ? OFFSET ?
    `, count, offset).Scan(&productResponse).Error; err != nil {
		return nil, err
	}

	return productResponse, nil
}

func (c *ProductRepository) Quantity(id uint64) (uint32, error) {
	var quantity uint32
	err := c.DB.Raw(`SELECT quantity FROM products WHERE id=?`, id).Scan(&quantity).Error
	if err != nil {
		return 0, err
	}
	return quantity, err
}

func (c *ProductRepository) FetchProduct(id uint64) (*pb.FetchProductResponse, error) {
	var product *pb.FetchProductResponse
	err := c.DB.Raw(`SELECT id,name,description,quantity,selling_price FROM products WHERE id=?`, id).Scan(&product).Error
	if err != nil {
		return &pb.FetchProductResponse{}, err
	}
	return product, err
}
