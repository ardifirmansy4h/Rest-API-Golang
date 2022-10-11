package products

import (
	"tugas/bussiness/products"

	"gorm.io/gorm"
)

type productRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) products.Repository {
	return &productRepository{
		conn: conn,
	}
}

func (pr *productRepository) GetAll() []products.Domain {
	var rec []Product

	pr.conn.Model(&rec).Preload("Order").Find(&rec)

	productDomain := []products.Domain{}

	for _, product := range rec {
		productDomain = append(productDomain, product.ToDomain())
	}

	return productDomain
}

func (pr *productRepository) GetByID(id string) products.Domain {
	var product Product
	pr.conn.Preload("Product").First(&product, "id = ?", id)

	return product.ToDomain()
}

func (pr *productRepository) Create(productDomain *products.Domain) products.Domain {
	rec := FromDomain(productDomain)

	res := pr.conn.Create(&rec)

	res.Last(&rec)
	return rec.ToDomain()
}

func (pr *productRepository) Update(id string, productDomain *products.Domain) products.Domain {
	var product products.Domain = pr.GetByID(id)

	UpdateProduct := FromDomain(&product)
	UpdateProduct.ItemCode = productDomain.ItemCode
	UpdateProduct.Description = product.Description
	UpdateProduct.Quantity = product.Quantity
	UpdateProduct.OrderID = product.OrderID

	pr.conn.Save(&UpdateProduct)

	return UpdateProduct.ToDomain()
}

func (pr *productRepository) Delete(id string) bool {
	var product products.Domain = pr.GetByID(id)

	deletedProduct := FromDomain(&product)
	res := pr.conn.Delete(&deletedProduct)

	if res.RowsAffected == 0 {
		return false
	}
	return true
}
