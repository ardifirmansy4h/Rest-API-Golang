package products

import (
	"tugas/bussiness/products"
)

type Product struct {
	ID          uint   `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

func (rec *Product) ToDomain() products.Domain {
	return products.Domain{
		ID:          rec.ID,
		ItemCode:    rec.ItemCode,
		Description: rec.Description,
		Quantity:    rec.Quantity,
		OrderID:     rec.OrderID,
	}
}

func FromDomain(domain *products.Domain) *Product {
	return &Product{
		ID:          domain.ID,
		ItemCode:    domain.ItemCode,
		Description: domain.Description,
		Quantity:    domain.Quantity,
		OrderID:     domain.OrderID,
	}
}
