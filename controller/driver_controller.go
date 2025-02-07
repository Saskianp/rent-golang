package controller

import (
	"testBeGo/models"
	"testBeGo/service"
	"testBeGo/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverController struct {
	Service *service.DriverService
}

func NewDriverController(service *service.DriverService) *DriverController {
	return &DriverController{Service: service}
}

func (c *DriverController) GetAll(ctx *gin.Context) {
	drivers, err := c.Service.GetAllDrivers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch drivers", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(drivers))
}

func (c *DriverController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	driver, err := c.Service.GetDriverByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Driver not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(driver))
}

func (c *DriverController) Create(ctx *gin.Context) {
	var driver models.Driver
	if err := ctx.ShouldBindJSON(&driver); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err := c.Service.CreateDriver(&driver)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create drivers", nil))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessfulResponse1("Driver created successfully"))
}

func (c *DriverController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var driver models.Driver
	if err := ctx.ShouldBindJSON(&driver); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err = c.Service.UpdateDriver(uint(id), &driver)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update driver", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Driver updated successfully"))
}

func (c *DriverController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	// Cek apakah driver dengan ID tersebut ada
	existingDriver, err := c.Service.GetDriverByID(uint(id))
	if err != nil || existingDriver.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Driver not found", nil))
		return
	}

	// Lanjutkan proses delete jika ID ditemukan
	err = c.Service.DeleteDriver(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to delete driver", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Driver deleted successfully"))
}
