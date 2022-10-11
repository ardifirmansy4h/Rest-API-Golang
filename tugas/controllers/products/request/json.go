package request

import "tugas/bussiness/products"

type Product struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

func (req *Product) ToDomain() *products.Domain {
	return &products.Domain{
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
		OrderID:     req.OrderID,
	}
}
