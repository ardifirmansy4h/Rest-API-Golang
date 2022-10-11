package products

type productUsecase struct {
	productRepository Repository
}

func NewProductUseCase(pr Repository) Usecase {
	return &productUsecase{
		productRepository: pr,
	}
}

func (pu *productUsecase) GetAll() []Domain {
	return pu.productRepository.GetAll()
}
func (pu *productUsecase) GetByID(id string) Domain {
	return pu.productRepository.GetByID(id)
}

func (pu *productUsecase) Create(orderDomain *Domain) Domain {
	return pu.productRepository.Create(orderDomain)
}

func (pu *productUsecase) Update(id string, orderDomain *Domain) Domain {
	return pu.productRepository.Update(id, orderDomain)
}
