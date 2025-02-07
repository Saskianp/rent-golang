package controller

import (
	"testBeGo/models"
	"testBeGo/service"
	"testBeGo/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MembershipController struct {
	Service *service.MembershipService
}

func NewMembershipController(service *service.MembershipService) *MembershipController {
	return &MembershipController{Service: service}
}

func (c *MembershipController) GetAll(ctx *gin.Context) {
	memberships, err := c.Service.GetAllMemberships()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to fetch Memberships", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(memberships))
}

func (c *MembershipController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	membership, err := c.Service.GetMembershipByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Membership not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1(membership))
}

func (c *MembershipController) Create(ctx *gin.Context) {
	var membership models.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err := c.Service.CreateMembership(&membership)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to create Memberships", nil))
		return
	}

	ctx.JSON(http.StatusCreated, helper.SuccessfulResponse1("Membership created successfully"))
}

func (c *MembershipController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var membership models.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid input data", nil))
		return
	}

	err = c.Service.UpdateMembership(uint(id), &membership)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to update Membership", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Membership updated successfully"))
}

func (c *MembershipController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.FailedResponse1(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	// Cek apakah Membership dengan ID tersebut ada
	existingMembership, err := c.Service.GetMembershipByID(uint(id))
	if err != nil || existingMembership.ID == 0 {
		ctx.JSON(http.StatusNotFound, helper.FailedResponse1(http.StatusNotFound, "Membership not found", nil))
		return
	}

	// Lanjutkan proses delete jika ID ditemukan
	err = c.Service.DeleteMembership(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.FailedResponse1(http.StatusInternalServerError, "Failed to delete Membership", nil))
		return
	}

	ctx.JSON(http.StatusOK, helper.SuccessfulResponse1("Membership deleted successfully"))
}
