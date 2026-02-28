package controllers

import (
	"net/http"

	"master-data-service-go/dto"
	"master-data-service-go/services"

	"github.com/gin-gonic/gin"
)

type ImplMasterJenisBarangController struct {
	service services.MasterJenisBarangService
}

func NewMasterJenisBarangController(service services.MasterJenisBarangService) *ImplMasterJenisBarangController {
	return &ImplMasterJenisBarangController{
		service: service,
	}
}

func (c *ImplMasterJenisBarangController) GetAll(ctx *gin.Context) {
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

func (c *ImplMasterJenisBarangController) GetByID(ctx *gin.Context) {
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

func (c *ImplMasterJenisBarangController) Create(ctx *gin.Context) {
	var payload dto.CreateMasterDataRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := c.service.Create(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func (c *ImplMasterJenisBarangController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload dto.UpdateMasterDataRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := c.service.Update(id, payload)
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

func (c *ImplMasterJenisBarangController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
func (c *ImplMasterJenisBarangController) SoftDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.SoftDelete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data soft deleted successfully",
	})
}
