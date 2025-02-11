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
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid data", err.Error()))
		return
	}

	// Simpan booking ke database
	createdBooking, err := c.service.CreateBooking(booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create booking", err.Error()))
		return
	}

	// Ambil ulang data booking dengan preloading customer, car, booking type, dan driver
	fullBooking, err := c.service.GetBookingByID(createdBooking.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to retrieve booking details", err.Error()))
		return
	}

	// Kirim respons dengan data lengkap
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(fullBooking))
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
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "ID is required", ""))
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

	// Hapus nilai kosong sebelum dikembalikan dalam respons
	if updatedBooking.Customer.ID == 0 {
		updatedBooking.Customer = models.Customer{}
	}
	if updatedBooking.Car.ID == 0 {
		updatedBooking.Car = models.Car{}
	}
	if updatedBooking.BookingType.ID == 0 {
		updatedBooking.BookingType = models.BookingType{}
	}
	if updatedBooking.Driver != nil && updatedBooking.Driver.ID == 0 {
		updatedBooking.Driver = nil
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
