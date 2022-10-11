package products

import (
	"net/http"
	"tugas/bussiness/products"
	"tugas/controllers"
	"tugas/controllers/products/request"
	"tugas/controllers/products/response"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUsecase products.Usecase
}

func NewProductController(productUC products.Usecase) *ProductController {
	return &ProductController{
		productUsecase: productUC,
	}
}

func (ctrl *ProductController) GetAllProduct(c echo.Context) error {
	productData := ctrl.productUsecase.GetAll()

	products := []response.Product{}

	for _, product := range productData {
		products = append(products, response.FromDomain(product))
	}

	res := controllers.NewResponse(c, http.StatusOK, "Succes", "All Products", products)

	return res
}

func (ctrl *ProductController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	product := ctrl.productUsecase.GetByID(id)

	if product.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "data kosong", "")
	}

	res := controllers.NewResponse(c, http.StatusOK, "success", "product found", response.FromDomain(product))
	return res
}

func (ctrl *ProductController) CreateProduct(c echo.Context) error {
	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "gagal tambah", "", "")
	}
	product := ctrl.productUsecase.Create(input.ToDomain())

	res := controllers.NewResponse(c, http.StatusCreated, "succes", "created product", response.FromDomain(product))

	return res
}

func (ctrl *ProductController) UpdateProduct(c echo.Context) error {
	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "gagal", "", "")
	}

	var id string = c.Param("id")

	product := ctrl.productUsecase.Update(id, input.ToDomain())

	res := controllers.NewResponse(c, http.StatusOK, "succes", "product update", response.FromDomain(product))

	return res
}
