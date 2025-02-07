package controller

import (
	"testBeGo/models"
	"testBeGo/service"
	"testBeGo/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverIncentiveController struct {
	Service *service.DriverIncentiveService
}

func NewDriverIncentiveController(service *service.DriverIncentiveService) *DriverIncentiveController {
	return &DriverIncentiveController{Service: service}
}

func (c *DriverIncentiveController) GetAll(ctx *gin.Context) {
	driverIncentives, err := c.Service.GetAllDriverIncentives()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch driverIncentives", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(driverIncentives))
}

func (c *DriverIncentiveController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	driverIncentive, err := c.Service.GetDriverIncentiveByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "DriverIncentive not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(driverIncentive))
}

func (c *DriverIncentiveController) Create(ctx *gin.Context) {
	var driverIncentive models.DriverIncentive
	if err := ctx.ShouldBindJSON(&driverIncentive); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err := c.Service.CreateDriverIncentive(&driverIncentive)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create driverIncentives", nil))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessfulResponse1("DriverIncentive created successfully"))
}

func (c *DriverIncentiveController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var driverIncentive models.DriverIncentive
	if err := ctx.ShouldBindJSON(&driverIncentive); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err = c.Service.UpdateDriverIncentive(uint(id), &driverIncentive)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update driverIncentive", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("DriverIncentive updated successfully"))
}

func (c *DriverIncentiveController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	// Cek apakah driverIncentive dengan ID tersebut ada
	existingDriverIncentive, err := c.Service.GetDriverIncentiveByID(uint(id))
	if err != nil || existingDriverIncentive.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "DriverIncentive not found", nil))
		return
	}

	// Lanjutkan proses delete jika ID ditemukan
	err = c.Service.DeleteDriverIncentive(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to delete driverIncentive", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("DriverIncentive deleted successfully"))
}
