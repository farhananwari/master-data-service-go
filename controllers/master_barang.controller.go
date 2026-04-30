package controllers

import (
	"net/http"

	"master-data-service-go/dto"
	"master-data-service-go/services"
	"master-data-service-go/utils"

	"github.com/gin-gonic/gin"
)

type ImplMasterBarangController struct {
	service services.MasterBarangService
}

func NewMasterBarangController(service services.MasterBarangService) *ImplMasterBarangController {
	return &ImplMasterBarangController{
		service: service,
	}
}

func (c *ImplMasterBarangController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, "success get data", data)
}

func (c *ImplMasterBarangController) GetBykodeBarang(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	data, err := c.service.GetByID(kodeBarang)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, "success get data", data)
}

func (c *ImplMasterBarangController) Create(ctx *gin.Context) {
	var payload dto.CreateMasterBarangRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errors := utils.FormatValidationError(err)
		utils.ValidationError(ctx, errors)
		return
	}

	data, err := c.service.Create(payload)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, http.StatusCreated, "berhasil create data", data)
}

func (c *ImplMasterBarangController) Update(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	var payload dto.UpdateMasterBarangRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := c.service.Update(kodeBarang, payload)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, http.StatusOK, "success update data", data)
}

func (c *ImplMasterBarangController) Delete(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	if err := c.service.Delete(kodeBarang); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dihapus",
	})
}

func (c *ImplMasterBarangController) SoftDelete(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	if err := c.service.SoftDelete(kodeBarang); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dihapus secara soft delete",
	})
}
