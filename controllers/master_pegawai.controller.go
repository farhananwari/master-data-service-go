package controllers

import (
	"net/http"

	"master-data-service-go/dto"
	"master-data-service-go/services"

	"github.com/gin-gonic/gin"
)

type ImplMasterPegawaiController struct {
	service services.MasterPegawaiService
}

func NewMasterPegawaiController(service services.MasterPegawaiService) *ImplMasterPegawaiController {
	return &ImplMasterPegawaiController{
		service: service,
	}
}

func (c *ImplMasterPegawaiController) GetAll(ctx *gin.Context) {
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

func (c *ImplMasterPegawaiController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := c.service.GetByID(id)
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

func (c *ImplMasterPegawaiController) Create(ctx *gin.Context) {
	var payload dto.CreateMasterPegawaiRequest

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

func (c *ImplMasterPegawaiController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var payload dto.UpdateMasterPegawaiRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := c.service.Update(id, payload)
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

func (c *ImplMasterPegawaiController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dihapus",
	})
}

func (c *ImplMasterPegawaiController) SoftDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.SoftDelete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dihapus secara soft delete",
	})
}
