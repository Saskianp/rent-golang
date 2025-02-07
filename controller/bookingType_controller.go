package controller

import (
	"testBeGo/models"
	"testBeGo/service"
	"testBeGo/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingTypeController struct {
	Service *service.BookingTypeService
}

func NewBookingTypeController(service *service.BookingTypeService) *BookingTypeController {
	return &BookingTypeController{Service: service}
}

func (c *BookingTypeController) GetAll(ctx *gin.Context) {
	bookingTypes, err := c.Service.GetAllBookingTypes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch Booking Types", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(bookingTypes))
}

func (c *BookingTypeController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	bookingType, err := c.Service.GetBookingTypeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "BookingType not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(bookingType))
}

func (c *BookingTypeController) Create(ctx *gin.Context) {
	var bookingType models.BookingType
	if err := ctx.ShouldBindJSON(&bookingType); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err := c.Service.CreateBookingType(&bookingType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create Booking Types", nil))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessfulResponse1("Booking Type created successfully"))
}

func (c *BookingTypeController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var bookingType models.BookingType
	if err := ctx.ShouldBindJSON(&bookingType); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err = c.Service.UpdateBookingType(uint(id), &bookingType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update Booking Type", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Booking Type updated successfully"))
}

func (c *BookingTypeController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	// Cek apakah BookingType dengan ID tersebut ada
	existingBookingType, err := c.Service.GetBookingTypeByID(uint(id))
	if err != nil || existingBookingType.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "BookingType not found", nil))
		return
	}

	// Lanjutkan proses delete jika ID ditemukan
	err = c.Service.DeleteBookingType(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to delete BookingType", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("BookingType deleted successfully"))
}
