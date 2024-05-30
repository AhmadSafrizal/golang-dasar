package main

import (
	"log"
	"net/http"

	handler "github.com/AhmadSafrizal/golang-dasar/tugas6/handlers"
	repository "github.com/AhmadSafrizal/golang-dasar/tugas6/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	engine := gin.New()

	engine.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "This is a test",
		})
	})

	myDb := "host=localhost user=postgres password=postgre dbname=prakerja port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(myDb), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	productRepo := &repository.ProductRepository{DB: db}

	productRepo.Migrate()

	productHandler := &handler.ProductHandler{Repository: productRepo}

	productGroup := engine.Group("/products")
	{
		productGroup.GET("", productHandler.GetGorm)
		productGroup.POST("", productHandler.CreateGrom)
		productGroup.PUT("/:id", productHandler.UpdateGorm)
		productGroup.DELETE("/:id", productHandler.DeleteGorm)
	}
	err = engine.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
