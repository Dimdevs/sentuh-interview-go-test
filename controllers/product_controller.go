package controllers

import (
	"net/http"
	"strconv"

	"go-test/config"
	"go-test/models"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create product"})
	}

	return c.JSON(http.StatusCreated, product)
}

func GetProduct(c echo.Context) error {
	var products []models.Product

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	query := config.DB.Preload("Category").Offset(offset).Limit(limit)

	if name := c.QueryParam("name"); name != "" {
		query = query.Where("products.name LIKE ?", "%"+name+"%")
	}
	if category := c.QueryParam("category"); category != "" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name LIKE ?", "%"+category+"%")
	}

	if err := query.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	var total int64
	countQuery := config.DB.Model(&models.Product{})
	if name := c.QueryParam("name"); name != "" {
		countQuery = countQuery.Where("name LIKE ?", "%"+name+"%")
	}
	if category := c.QueryParam("category"); category != "" {
		countQuery = countQuery.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name LIKE ?", "%"+category+"%")
	}
	if err := countQuery.Count(&total).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)
	return c.JSON(http.StatusOK, echo.Map{
		"products":    products,
		"total":       total,
		"page":        page,
		"total_pages": totalPages,
	})
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid product ID"})
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := config.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update product"})
	}

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid product ID"})
	}

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete product"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
