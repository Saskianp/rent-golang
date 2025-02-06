package controller

import (
	"testBeGo/models"
	"testBeGo/service"
	"testBeGo/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service *service.CustomerService
}

func NewCustomerController(service *service.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

func (c *CustomerController) GetAll(ctx *gin.Context) {
	customers, err := c.Service.GetAllCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch customers", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(customers))
}

func (c *CustomerController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	customer, err := c.Service.GetCustomerByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Customer not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(customer))
}

func (c *CustomerController) Create(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err := c.Service.CreateCustomer(&customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create customer", nil))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessfulResponse1("Customer created successfully"))
}

func (c *CustomerController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err = c.Service.UpdateCustomer(uint(id), &customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update customer", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Customer updated successfully"))
}

func (c *CustomerController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	// Cek apakah customer dengan ID tersebut ada
	existingCustomer, err := c.Service.GetCustomerByID(uint(id))
	if err != nil || existingCustomer.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Customer not found", nil))
		return
	}

	// Lanjutkan proses delete jika ID ditemukan
	err = c.Service.DeleteCustomer(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to delete customer", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Customer deleted successfully"))
}
