package response

import (
	"time"
	"tugas/bussiness/orders"
	"tugas/drivers/mysql/products"
)

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Product      []products.Product
}

func FromDomain(domain orders.Domain) Order {
	return Order{
		ID:           domain.ID,
		CustomerName: domain.CustomerName,
		OrderedAt:    domain.OrderedAt,
		Product:      domain.Product,
	}
}
