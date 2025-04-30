package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pinduoduo-back/database"
	"pinduoduo-back/order-service/models"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupProductTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/products", CreateProduct)
	router.GET("/products", GetProducts)
	router.GET("/products/:id", GetProduct)
	router.PUT("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", DeleteProduct)
	return router
}

func initProductTestDB() {
	database.ConnectTestDB()
	database.DB = database.TestDB
	database.DB.Exec("DELETE FROM products")
}

func TestCreateProduct(t *testing.T) {
	initProductTestDB()
	router := setupProductTestRouter()

	product := models.Product{Name: "Test Product", Price: 19.99}
	body, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var created models.Product
	err := json.Unmarshal(w.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, created.Name)
	assert.Equal(t, product.Price, created.Price)
}

func TestGetProducts(t *testing.T) {
	initProductTestDB()
	database.DB.Create(&models.Product{Name: "Sample Product", Price: 10.0})
	router := setupProductTestRouter()

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var products []models.Product
	err := json.Unmarshal(w.Body.Bytes(), &products)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(products), 1)
}

func TestGetProduct(t *testing.T) {
	initProductTestDB()
	product := models.Product{Name: "Unique Product", Price: 15.5}
	database.DB.Create(&product)

	router := setupProductTestRouter()
	req, _ := http.NewRequest("GET", "/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var fetched models.Product
	err := json.Unmarshal(w.Body.Bytes(), &fetched)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, fetched.Name)
}

func TestUpdateProduct(t *testing.T) {
	initProductTestDB()
	product := models.Product{Name: "Old Name", Price: 10.0}
	database.DB.Create(&product)

	updated := models.Product{Name: "New Name", Price: 20.0}
	body, _ := json.Marshal(updated)

	router := setupProductTestRouter()
	req, _ := http.NewRequest("PUT", "/products/"+strconv.Itoa(int(product.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetched models.Product
	err := json.Unmarshal(w.Body.Bytes(), &fetched)
	assert.NoError(t, err)
	assert.Equal(t, updated.Name, fetched.Name)
}

func TestDeleteProduct(t *testing.T) {
	initProductTestDB()
	product := models.Product{Name: "To Delete", Price: 5.0}
	database.DB.Create(&product)

	router := setupProductTestRouter()
	req, _ := http.NewRequest("DELETE", "/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что продукт удален
	var result models.Product
	err := database.DB.First(&result, product.ID).Error
	assert.Error(t, err)
}
