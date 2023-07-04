package handlers

import (
	dto "dumbmerch/dto/result"
	usersdto "dumbmerch/dto/user"
	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repository.UserRepository
}

func UserHandler(userRepository repository.UserRepository) *handler {
	return &handler{userRepository}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(user)})
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		// Email:    u.Email,
		// Password: u.Password,
	}
}
