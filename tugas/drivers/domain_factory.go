package drivers

import (
	orderDomain "tugas/bussiness/orders"
	productDomain "tugas/bussiness/products"
	orderDB "tugas/drivers/mysql/orders"
	productDB "tugas/drivers/mysql/products"

	"gorm.io/gorm"
)

func NewProductRepository(conn *gorm.DB) productDomain.Repository {
	return productDB.NewMySQLRepository(conn)
}
func NewOrderRepository(conn *gorm.DB) orderDomain.Repository {
	return orderDB.NewMySQLRepository(conn)
}
