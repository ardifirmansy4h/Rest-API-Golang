package orders

import (
	"tugas/bussiness/orders"

	"gorm.io/gorm"
)

type orderRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) orders.Repository {
	return &orderRepository{
		conn: conn,
	}
}

func (or *orderRepository) GetAll() []orders.Domain {
	var rec []Order
	or.conn.Model(&rec).Preload("Product").Find(&rec)

	orderDomain := []orders.Domain{}
	for _, order := range rec {
		orderDomain = append(orderDomain, order.ToDomain())
	}
	return orderDomain
}

func (or *orderRepository) GetByID(id string) orders.Domain {
	var order Order
	or.conn.Preload("Product").First(&order, "id = ?", id)

	return order.ToDomain()
}

func (or *orderRepository) Create(orderDomain *orders.Domain) orders.Domain {
	rec := FromDomain(orderDomain)

	res := or.conn.Create(&rec)

	res.Last(&rec)
	return rec.ToDomain()
}

func (or *orderRepository) Update(id string, orderDomain *orders.Domain) orders.Domain {
	var order orders.Domain = or.GetByID(id)

	UpdateOrder := FromDomain(&order)
	UpdateOrder.CustomerName = orderDomain.CustomerName
	UpdateOrder.OrderedAt = order.OrderedAt

	or.conn.Save(&UpdateOrder)

	return UpdateOrder.ToDomain()
}

func (or *orderRepository) Delete(id string) bool {
	var order orders.Domain = or.GetByID(id)

	deletedOrder := FromDomain(&order)
	res := or.conn.Delete(&deletedOrder)

	if res.RowsAffected == 0 {
		return false
	}
	return true
}
