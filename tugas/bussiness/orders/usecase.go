package orders

type orderUsecase struct {
	orderRepository Repository
}

func NewOrderUsecase(or Repository) Usecase {
	return &orderUsecase{
		orderRepository: or,
	}
}

func (ou *orderUsecase) GetAll() []Domain {
	return ou.orderRepository.GetAll()
}
func (ou *orderUsecase) GetByID(id string) Domain {
	return ou.orderRepository.GetByID(id)
}

func (ou *orderUsecase) Create(orderDomain *Domain) Domain {
	return ou.orderRepository.Create(orderDomain)
}

func (ou *orderUsecase) Update(id string, orderDomain *Domain) Domain {
	return ou.orderRepository.Update(id, orderDomain)
}

func (ou *orderUsecase) Delete(id string) bool {
	return ou.orderRepository.Delete(id)
}
