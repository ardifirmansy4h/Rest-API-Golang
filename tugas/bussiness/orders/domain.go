package orders

import (
	"time"
	"tugas/drivers/mysql/products"
)

type Domain struct {
	ID           uint
	CustomerName string
	Product      []products.Product
	OrderedAt    time.Time
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(orderDomain *Domain) Domain
	Update(id string, orderDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(orderDomain *Domain) Domain
	Update(id string, orderDomain *Domain) Domain
	Delete(id string) bool
}
