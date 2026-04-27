package controllers

import (
	"net/http"

	"master-data-service-go/dto"
	"master-data-service-go/services"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (c *ImplMasterBarangController) GetBykodeBarang(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	data, err := c.service.GetByID(kodeBarang)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (c *ImplMasterBarangController) Create(ctx *gin.Context) {
	var payload dto.CreateMasterBarangRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := c.service.Create(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func (c *ImplMasterBarangController) Update(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	var payload dto.UpdateMasterBarangRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := c.service.Update(kodeBarang, payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (c *ImplMasterBarangController) Delete(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	if err := c.service.Delete(kodeBarang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dihapus",
	})
}

func (c *ImplMasterBarangController) SoftDelete(ctx *gin.Context) {
	kodeBarang := ctx.Param("kodeBarang")

	if err := c.service.SoftDelete(kodeBarang); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dihapus secara soft delete",
	})
}
