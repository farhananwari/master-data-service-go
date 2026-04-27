package routes

import (
	"master-data-service-go/controllers"
	"master-data-service-go/repositories"
	"master-data-service-go/services"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	// Initialize repositories
	masterBarangRepo := repositories.NewMasterBarangRepository(db)
	masterJenisBarangRepo := repositories.NewMasterJenisBarangRepository(db)
	masterSatuanBrangRepo := repositories.NewMasterSatuanBarangRepository(db)
	masterLokasiRepo := repositories.NewMasterLokasiRepository(db)
	masterPegawaiRepo := repositories.NewMasterPegawaiRepository(db)

	// Initialize services
	masterBarangService := services.NewMasterBarangService(masterBarangRepo)
	masterJenisBarangService := services.NewMasterJenisBarangService(masterJenisBarangRepo)
	masterSatuanBarangService := services.NewMasterSatuanBarangService(masterSatuanBrangRepo)
	masterLokasiService := services.NewMasterLokasiService(masterLokasiRepo)
	masterPegawaiService := services.NewMasterPegawaiService(masterPegawaiRepo)

	// Initialize controller
	masterBarangController := controllers.NewMasterBarangController(masterBarangService)
	masterJenisBarangController := controllers.NewMasterJenisBarangController(masterJenisBarangService)
	masterSatuanBarangController := controllers.NewMasterSatuanBarangController(masterSatuanBarangService)
	masterLokasiController := controllers.NewMasterLokasiController(masterLokasiService)
	masterPegawaiController := controllers.NewMasterPegawaiController(masterPegawaiService)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "master-data-service-go",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Root endpoint
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "master-data Service API v1",
				"version": "1.0.0",
			})
		})
	}

	routeMasterBarang := router.Group("/api/master-barang")
	{
		routeMasterBarang.GET("/", masterBarangController.GetAll)
		routeMasterBarang.GET("/:kodeBarang", masterBarangController.GetBykodeBarang)
		routeMasterBarang.POST("/", masterBarangController.Create)
		routeMasterBarang.PUT("/:kodeBarang", masterBarangController.Update)
		//routeMasterBarang.DELETE("/:id", masterBarangController.Delete)
		routeMasterBarang.DELETE("/:kodeBarang", masterBarangController.SoftDelete)
	}

	routeMasterJenisBarang := router.Group("/api/master-jenis-barang")
	{
		routeMasterJenisBarang.GET("/", masterJenisBarangController.GetAll)
		routeMasterJenisBarang.GET("/:id", masterJenisBarangController.GetByID)
		routeMasterJenisBarang.POST("/", masterJenisBarangController.Create)
		routeMasterJenisBarang.PUT("/:id", masterJenisBarangController.Update)
		routeMasterJenisBarang.DELETE("/:id", masterJenisBarangController.SoftDelete)
	}

	routeMasterSatuanBarang := router.Group("/api/master-satuan-barang")
	{
		routeMasterSatuanBarang.GET("/", masterSatuanBarangController.GetAll)
		routeMasterSatuanBarang.GET("/:id", masterSatuanBarangController.GetByID)
		routeMasterSatuanBarang.POST("/", masterSatuanBarangController.Create)
		routeMasterSatuanBarang.PUT("/:id", masterSatuanBarangController.Update)
		//routeMasterSatuanBarang.DELETE("/:id", masterSatuanBarangController.Delete)
		routeMasterSatuanBarang.DELETE("/:id", masterSatuanBarangController.SoftDelete)
	}

	routeMasterLokasi := router.Group("/api/master-lokasi")
	{
		routeMasterLokasi.GET("/", masterLokasiController.GetAll)
		routeMasterLokasi.GET("/:id", masterLokasiController.GetByID)
		routeMasterLokasi.POST("/", masterLokasiController.Create)
		routeMasterLokasi.PUT("/:id", masterLokasiController.Update)
		routeMasterLokasi.DELETE("/:id", masterLokasiController.SoftDelete)
	}

	routeMasterPegawai := router.Group("/api/master-pegawai")
	{
		routeMasterPegawai.GET("/", masterPegawaiController.GetAll)
		routeMasterPegawai.GET("/:id", masterPegawaiController.GetByID)
		routeMasterPegawai.POST("/", masterPegawaiController.Create)
		routeMasterPegawai.PUT("/:id", masterPegawaiController.Update)
		//routeMasterPegawai.DELETE("/:id", masterPegawaiController.Delete)
		routeMasterPegawai.DELETE("/:id", masterPegawaiController.SoftDelete)
	}
}
