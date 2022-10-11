package route

import (
	"tugas/controllers/orders"
	"tugas/controllers/products"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	OrderController   orders.OrderController
	ProductController products.ProductController
}

func (cl *ControllerList) SetRoutes(e *echo.Echo) {
	e.GET("/orders", cl.OrderController.GetAllOrder)
	e.GET("/orders/:id", cl.OrderController.GetByID)
	e.POST("/orders", cl.OrderController.CreateOrder)
	e.DELETE("/orders/:id", cl.OrderController.DeleteOrder)
	e.PUT("/orders/:id", cl.OrderController.UpdateOrder)
	e.GET("/products", cl.ProductController.GetAllProduct)
	e.GET("/products/:id", cl.ProductController.GetByID)
	e.POST("/products", cl.ProductController.CreateProduct)
	e.PUT("/products/:id", cl.ProductController.UpdateProduct)
}
