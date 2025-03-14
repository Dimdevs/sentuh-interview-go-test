package controllers

import (
	"encoding/xml"
	"go-test/config"
	"go-test/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	XMLName  xml.Name `xml:"CreateUserRequest"`
	Name     string   `xml:"name"`
	Email    string   `xml:"email"`
	Password string   `xml:"password"`
}

type UpdateUserRequest struct {
	XMLName xml.Name `xml:"UpdateUserRequest"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

type ErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Message string   `xml:"message"`
}

type UserResponse struct {
	XMLName xml.Name `xml:"user"`
	ID      uint     `xml:"id"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

type StatusResponse struct {
	XMLName xml.Name      `xml:"Response"`
	Status  string        `xml:"status"`
	User    *UserResponse `xml:"user,omitempty"`
}

type UserListResponse struct {
	XMLName    xml.Name       `xml:"GetUsersResponse"`
	Users      []UserResponse `xml:"users>user"`
	Pagination struct {
		Page  int   `xml:"page"`
		Limit int   `xml:"limit"`
		Total int64 `xml:"total"`
	} `xml:"pagination"`
}

func toUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func SoapCreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := xml.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.XML(http.StatusBadRequest, ErrorResponse{Message: "Invalid XML"})
	}

	user := models.User{Name: req.Name, Email: req.Email, Password: req.Password}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ErrorResponse{Message: "Gagal menyimpan user"})
	}

	respData := toUserResponse(user)
	resp := StatusResponse{
		Status: "success",
		User:   &respData,
	}
	return c.XML(http.StatusCreated, resp)
}

func SoapGetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.XML(http.StatusNotFound, ErrorResponse{Message: "User tidak ditemukan"})
	}
	return c.XML(http.StatusOK, toUserResponse(user))
}

func SoapGetUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var users []models.User
	if err := config.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ErrorResponse{Message: "Gagal mengambil data"})
	}

	var total int64
	config.DB.Model(&models.User{}).Count(&total)

	var userResponses []UserResponse
	for _, u := range users {
		userResponses = append(userResponses, toUserResponse(u))
	}

	resp := UserListResponse{
		Users: userResponses,
	}
	resp.Pagination.Page = page
	resp.Pagination.Limit = limit
	resp.Pagination.Total = total

	return c.XML(http.StatusOK, resp)
}

func SoapUpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.XML(http.StatusNotFound, ErrorResponse{Message: "User tidak ditemukan"})
	}

	var req UpdateUserRequest
	if err := xml.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.XML(http.StatusBadRequest, ErrorResponse{Message: "Invalid XML"})
	}

	user.Name = req.Name
	user.Email = req.Email

	if err := config.DB.Save(&user).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ErrorResponse{Message: "Gagal update user"})
	}

	respData := toUserResponse(user)
	resp := StatusResponse{
		Status: "success",
		User:   &respData,
	}
	return c.XML(http.StatusOK, resp)
}

func SoapDeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.XML(http.StatusInternalServerError, ErrorResponse{Message: "Gagal hapus user"})
	}

	return c.XML(http.StatusOK, StatusResponse{Status: "success"})
}
