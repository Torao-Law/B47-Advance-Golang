package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	repo := repository.MakeRepository(mysql.DB)
	h := handlers.ProductHandler(repo)

	e.GET("/product/:id", h.GetProduct)
	e.GET("/products", h.FindProducts)
	e.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
}
