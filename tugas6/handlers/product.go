package handler

import (
	"net/http"
	"strconv"

	model "github.com/AhmadSafrizal/golang-dasar/tugas6/models"
	repository "github.com/AhmadSafrizal/golang-dasar/tugas6/repositories"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repository *repository.ProductRepository
}

func (prod *ProductHandler) GetGorm(ctx *gin.Context) {
	products, err := prod.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
	}
	ctx.JSON(http.StatusOK, products)
}

func (prod *ProductHandler) CreateGrom(ctx *gin.Context) {
	product := &model.Product{}

	if err := ctx.Bind(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "invalid body request",
		})
		return
	}

	err := prod.Repository.Create(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

func (prod *ProductHandler) UpdateGorm(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product := &model.Product{}
	if err := ctx.Bind(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid body request"})
		return
	}

	product.ID = int(id)
	if err := prod.Repository.Update(uint(id), product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (prod *ProductHandler) DeleteGorm(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := prod.Repository.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
