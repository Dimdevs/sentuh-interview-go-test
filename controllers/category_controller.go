package controllers

import (
	"encoding/xml"
	"net/http"
	"strconv"

	"go-test/config"
	"go-test/models"

	"github.com/labstack/echo/v4"
)

type CategoryRequest struct {
	XMLName xml.Name `xml:"CategoryRequest"`
	Name    string   `xml:"name"`
}

type CategoryResponse struct {
	XMLName xml.Name `xml:"category"`
	ID      uint     `xml:"id"`
	Name    string   `xml:"name"`
}

type CategoryListResponse struct {
	XMLName    xml.Name           `xml:"GetCategoriesResponse"`
	Categories []CategoryResponse `xml:"categories>category"`
	Pagination struct {
		Page  int   `xml:"page"`
		Limit int   `xml:"limit"`
		Total int64 `xml:"total"`
	} `xml:"pagination"`
}

type CategoryStatusResponse struct {
	XMLName  xml.Name          `xml:"CategoryStatusResponse"`
	Status   string            `xml:"status"`
	Category *CategoryResponse `xml:"category,omitempty"`
}

type CategoryErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Message string   `xml:"message"`
}

func toCategoryResponse(c models.Category) CategoryResponse {
	return CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}

func SoapCreateCategory(c echo.Context) error {
	var req CategoryRequest
	if err := xml.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.XML(http.StatusBadRequest, CategoryErrorResponse{Message: "Invalid XML"})
	}
	category := models.Category{Name: req.Name}
	if err := config.DB.Create(&category).Error; err != nil {
		return c.XML(http.StatusInternalServerError, CategoryErrorResponse{Message: "Failed to create category"})
	}

	respData := toCategoryResponse(category)
	resp := CategoryStatusResponse{
		Status:   "success",
		Category: &respData,
	}
	return c.XML(http.StatusCreated, resp)
}

func SoapGetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.XML(http.StatusNotFound, CategoryErrorResponse{Message: "Category not found"})
	}
	return c.XML(http.StatusOK, toCategoryResponse(category))
}

func SoapGetCategories(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var categories []models.Category
	if err := config.DB.Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return c.XML(http.StatusInternalServerError, CategoryErrorResponse{Message: "Failed to fetch categories"})
	}

	var total int64
	config.DB.Model(&models.Category{}).Count(&total)

	var responseList []CategoryResponse
	for _, cat := range categories {
		responseList = append(responseList, toCategoryResponse(cat))
	}

	resp := CategoryListResponse{
		Categories: responseList,
	}
	resp.Pagination.Page = page
	resp.Pagination.Limit = limit
	resp.Pagination.Total = total

	return c.XML(http.StatusOK, resp)
}

func SoapUpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.XML(http.StatusNotFound, CategoryErrorResponse{Message: "Category not found"})
	}
	var req CategoryRequest
	if err := xml.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.XML(http.StatusBadRequest, CategoryErrorResponse{Message: "Invalid XML"})
	}
	category.Name = req.Name
	if err := config.DB.Save(&category).Error; err != nil {
		return c.XML(http.StatusInternalServerError, CategoryErrorResponse{Message: "Failed to update category"})
	}
	respData := toCategoryResponse(category)
	resp := CategoryStatusResponse{
		Status:   "success",
		Category: &respData,
	}
	return c.XML(http.StatusOK, resp)
}

func SoapDeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.XML(http.StatusInternalServerError, CategoryErrorResponse{Message: "Failed to delete category"})
	}
	return c.XML(http.StatusOK, CategoryStatusResponse{Status: "success"})
}
