package controller

import (
	"net/http"
	"strconv"
	"testBeGo/helper"
	"testBeGo/models"
	"testBeGo/service"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	service service.BookingService
}

func NewBookingController(service service.BookingService) *BookingController {
	return &BookingController{service}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
	var booking models.Booking
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid data", nil))
		return
	}

	createdBooking, err := c.service.CreateBooking(booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create booking", nil))
		return
	}

	// Ensure that the response includes the Customer and Car data
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(createdBooking))
}

func (c *BookingController) GetAllBookings(ctx *gin.Context) {
	bookings, err := c.service.GetAllBookings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch bookings", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(bookings))
}

func (c *BookingController) GetBookingByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	if idStr == "" {
		idStr = ctx.Query("id")
	}

	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "ID is required", nil))
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", err.Error()))
		return
	}

	booking, err := c.service.GetBookingByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Booking not found", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(booking))
}

func (c *BookingController) UpdateBooking(ctx *gin.Context) {
	var booking models.Booking

	// Ambil `id` dari path atau query parameter
	idStr := ctx.Param("id")
	if idStr == "" {
		idStr = ctx.Query("id")
	}

	// Jika ID kosong, return error
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "ID is required", nil))
		return
	}

	// Konversi `id` ke uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", err.Error()))
		return
	}

	// Bind JSON ke struct
	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid data format", err.Error()))
		return
	}

	// Update booking
	updatedBooking, err := c.service.UpdateBooking(uint(id), booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update booking", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(updatedBooking))
}

func (c *BookingController) DeleteBooking(ctx *gin.Context) {
	// Ambil `id` dari path atau query parameter
	idStr := ctx.Param("id")
	if idStr == "" {
		idStr = ctx.Query("id")
	}

	// Jika ID kosong, return error
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "ID is required", nil))
		return
	}

	// Konversi `id` ke uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", err.Error()))
		return
	}

	// Hapus booking
	err = c.service.DeleteBooking(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Booking not found", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Booking deleted successfully"))
}
