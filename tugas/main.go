package main

import (
	route "tugas/app/routes"
	orderUseCase "tugas/bussiness/orders"
	productUseCase "tugas/bussiness/products"
	orderController "tugas/controllers/orders"
	productController "tugas/controllers/products"
	drivers "tugas/drivers"
	dbDriver "tugas/drivers/mysql"
	"tugas/helper"

	"github.com/labstack/echo/v4"
)

func main() {

	configDB := dbDriver.ConfigDB{
		DB_USERNAME: helper.GetConfig("DB_USERNAME"),
		DB_PASSWORD: helper.GetConfig("DB_PASSWORD"),
		DB_NAME:     helper.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()
	dbDriver.DBMigrate(db)
	e := echo.New()

	orderRepo := drivers.NewOrderRepository(db)
	orderUsecase := orderUseCase.NewOrderUsecase(orderRepo)
	orderCtrl := orderController.NewOrdersController(orderUsecase)

	productRepo := drivers.NewProductRepository(db)
	productUseCase := productUseCase.NewProductUseCase(productRepo)
	productCtrl := productController.NewProductController(productUseCase)

	routesInit := route.ControllerList{
		OrderController:   *orderCtrl,
		ProductController: *productCtrl,
	}
	routesInit.SetRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
