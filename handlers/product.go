package handlers

import (
	productdto "dumbmerch/dto/product"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repository.ProductRepository
}

func ProductHandler(productRepository repository.ProductRepository) *handlerProduct {
	return &handlerProduct{productRepository}
}

func (h *handlerProduct) FindProducts(c echo.Context) error {
	users, err := h.ProductRepository.FindProducts()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertProductResponse(product)})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	request := new(productdto.ProductRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	data := models.Product{
		Name:      request.Name,
		Desc:      request.Desc,
		Price:     request.Price,
		Image:     request.Image,
		Qty:       request.Qty,
		UserID:    request.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := h.ProductRepository.CreateProduct(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func convertProductResponse(u models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		Name:   u.Name,
		Desc:   u.Desc,
		Price:  u.Price,
		Image:  u.Image,
		Qty:    u.Qty,
		UserID: u.UserID,
	}
}
