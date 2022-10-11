package products

type Domain struct {
	ID          uint
	ItemCode    string
	Description string
	Quantity    int
	OrderID     uint
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(orderDomain *Domain) Domain
	Update(id string, orderDomain *Domain) Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(orderDomain *Domain) Domain
	Update(id string, orderDomain *Domain) Domain
}
