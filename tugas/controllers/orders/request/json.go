package request

import "tugas/bussiness/orders"

type Order struct {
	CustomerName string `json:"customer_name"`
}

func (req *Order) ToDomain() *orders.Domain {
	return &orders.Domain{
		CustomerName: req.CustomerName,
	}
}
