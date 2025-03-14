package controllers

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"go-test/config"
	"go-test/models"

	"github.com/labstack/echo/v4"
)

type ProductRequest struct {
	XMLName    xml.Name `xml:"ProductRequest"`
	Name       string   `xml:"name"`
	Price      float64  `xml:"price"`
	CategoryID uint     `xml:"category_id"`
}

type ProductResponse struct {
	XMLName    xml.Name `xml:"product"`
	ID         uint     `xml:"id"`
	Name       string   `xml:"name"`
	Price      float64  `xml:"price"`
	CategoryID uint     `xml:"category_id"`
	Category   string   `xml:"category"`
}

type ProductListResponse struct {
	XMLName    xml.Name          `xml:"GetProductsResponse"`
	Products   []ProductResponse `xml:"products>product"`
	Pagination struct {
		Page  int   `xml:"page"`
		Limit int   `xml:"limit"`
		Total int64 `xml:"total"`
	} `xml:"pagination"`
}

type ProductStatusResponse struct {
	XMLName xml.Name         `xml:"ProductStatusResponse"`
	Status  string           `xml:"status"`
	Product *ProductResponse `xml:"product,omitempty"`
}

type ProductErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Message string   `xml:"message"`
}

func toProductResponse(p models.Product) ProductResponse {
	categoryName := ""
	categoryID := uint(0)

	if p.CategoryID != nil {
		categoryID = *p.CategoryID
	}
	if p.Category != nil {
		categoryName = p.Category.Name
	}

	return ProductResponse{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		CategoryID: categoryID,
		Category:   categoryName,
	}
}

func SoapCreateProduct(c echo.Context) error {
	var req ProductRequest
	if err := xml.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.XML(http.StatusBadRequest, ProductErrorResponse{Message: "Invalid XML"})
	}

	if req.Price < 0 {
		return c.XML(http.StatusBadRequest, ProductErrorResponse{Message: "Price cannot be negative"})
	}

	var category models.Category
	if err := config.DB.First(&category, req.CategoryID).Error; err != nil {
		return c.XML(http.StatusBadRequest, ProductErrorResponse{Message: "Category ID not found"})
	}

	product := models.Product{
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: &req.CategoryID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ProductErrorResponse{Message: "Failed to create product"})
	}

	config.DB.Preload("Category").First(&product)

	respData := toProductResponse(product)
	return c.XML(http.StatusCreated, ProductStatusResponse{
		Status:  "success",
		Product: &respData,
	})
}

func SoapGetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := config.DB.Preload("Category").First(&product, id).Error; err != nil {
		return c.XML(http.StatusNotFound, ProductErrorResponse{Message: "Product not found"})
	}

	return c.XML(http.StatusOK, toProductResponse(product))
}

func SoapGetProducts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var products []models.Product
	query := config.DB.Offset(offset).Limit(limit).Preload("Category")

	if name := c.QueryParam("name"); name != "" {
		query = query.Where("products.name LIKE ?", "%"+name+"%")
	}
	if category := c.QueryParam("category"); category != "" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.name LIKE ?", "%"+category+"%")
	}

	if err := query.Find(&products).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ProductErrorResponse{Message: "Failed to fetch products"})
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
	countQuery.Count(&total)

	var list []ProductResponse
	for _, p := range products {
		list = append(list, toProductResponse(p))
	}

	resp := ProductListResponse{Products: list}
	resp.Pagination.Page = page
	resp.Pagination.Limit = limit
	resp.Pagination.Total = total

	return c.XML(http.StatusOK, resp)
}

func SoapUpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		return c.XML(http.StatusNotFound, ProductErrorResponse{Message: "Product not found"})
	}

	var req ProductRequest
	if err := xml.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.XML(http.StatusBadRequest, ProductErrorResponse{Message: "Invalid XML"})
	}

	if req.Price < 0 {
		return c.XML(http.StatusBadRequest, ProductErrorResponse{Message: "Price cannot be negative"})
	}

	var category models.Category
	if err := config.DB.First(&category, req.CategoryID).Error; err != nil {
		return c.XML(http.StatusBadRequest, ProductErrorResponse{Message: "Category ID not found"})
	}

	product.Name = req.Name
	product.Price = req.Price
	product.CategoryID = &req.CategoryID

	if err := config.DB.Save(&product).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ProductErrorResponse{Message: "Failed to update product"})
	}

	config.DB.Preload("Category").First(&product)

	respData := toProductResponse(product)
	return c.XML(http.StatusOK, ProductStatusResponse{
		Status:  "success",
		Product: &respData,
	})
}

func SoapDeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ProductErrorResponse{Message: "Failed to delete product"})
	}

	return c.XML(http.StatusOK, ProductStatusResponse{Status: "success"})
}
