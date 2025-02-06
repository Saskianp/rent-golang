package controller

import (
	"testBeGo/helper"
	"testBeGo/models"
	"testBeGo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	Service *service.CarService
}

func NewCarController(service *service.CarService) *CarController {
	return &CarController{Service: service}
}

func (c *CarController) GetAll(ctx *gin.Context) {
	cars, err := c.Service.GetAllCars()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch cars", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(cars))
}

func (c *CarController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	car, err := c.Service.GetCarByID(uint(id))
	if err != nil || car.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Car not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(car))
}

func (c *CarController) Create(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err := c.Service.CreateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create car", nil))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessfulResponse1("Car created successfully"))
}

func (c *CarController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err = c.Service.UpdateCar(uint(id), &car)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update car", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Car updated successfully"))
}

func (c *CarController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	existingCar, err := c.Service.GetCarByID(uint(id))
	if err != nil || existingCar.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Car not found", nil))
		return
	}

	err = c.Service.DeleteCar(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to delete car", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Car deleted successfully"))
}
