package orders

import (
	"net/http"
	"tugas/bussiness/orders"
	"tugas/controllers"

	"tugas/controllers/orders/request"
	"tugas/controllers/orders/response"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderUseCase orders.Usecase
}

func NewOrdersController(orderUC orders.Usecase) *OrderController {
	return &OrderController{
		orderUseCase: orderUC,
	}
}

func (ctrl *OrderController) GetAllOrder(c echo.Context) error {
	ordersData := ctrl.orderUseCase.GetAll()

	orders := []response.Order{}

	for _, order := range ordersData {
		orders = append(orders, response.FromDomain(order))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all orders", orders)
}

func (ctrl *OrderController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	order := ctrl.orderUseCase.GetByID(id)

	if order.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "orders not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "orders found", response.FromDomain(order))
}

func (ctrl *OrderController) CreateOrder(c echo.Context) error {
	input := request.Order{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "gagal", "", "")
	}

	order := ctrl.orderUseCase.Create(input.ToDomain())

	res := controllers.NewResponse(c, http.StatusCreated, "succes", "orders created", response.FromDomain(order))

	return res
}

func (ctrl *OrderController) UpdateOrder(c echo.Context) error {
	input := request.Order{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "gagal", "", "")
	}

	var id string = c.Param("id")

	order := ctrl.orderUseCase.Update(id, input.ToDomain())

	res := controllers.NewResponse(c, http.StatusOK, "succes", "order update", response.FromDomain(order))

	return res
}

func (ctrl *OrderController) DeleteOrder(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := ctrl.orderUseCase.Delete(id)

	if !isDeleted {
		return controllers.NewResponse(c, http.StatusNotFound, "gagal", "order not found", "")
	}

	res := controllers.NewResponse(c, http.StatusOK, "succes", "order deleted", "")

	return res
}
