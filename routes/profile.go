package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func ProfileRoute(e *echo.Group) {
	repo := repository.MakeRepository(mysql.DB)
	h := handlers.ProfileHandler(repo)

	e.GET("/profile/:id", h.GetProfile)
}
