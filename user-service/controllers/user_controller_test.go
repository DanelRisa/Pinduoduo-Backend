package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"pinduoduo-back/database"
	"pinduoduo-back/user-service/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func setupUserTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
	return r
}

func initUserTestDB() {
	database.ConnectTestDB()
	database.TestDB.Exec("DELETE FROM users")
	database.DB = database.TestDB
}

func TestGetUsers(t *testing.T) {
	initUserTestDB()
	router := setupUserTestRouter()

	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	database.DB.Create(&user)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(users), 1)
}

func TestGetUser(t *testing.T) {
	initUserTestDB()
	router := setupUserTestRouter()

	user := models.User{
		Username: "findme",
		Email:    "find@example.com",
		Password: "securepass",
	}
	database.DB.Create(&user)

	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var found models.User
	err := json.Unmarshal(w.Body.Bytes(), &found)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, found.ID)
}

func TestUpdateUser(t *testing.T) {
	initUserTestDB()
	router := setupUserTestRouter()

	user := models.User{
		Username: "updatable",
		Email:    "old@mail.com",
		Password: "pass",
	}
	database.DB.Create(&user)

	updatedData := models.User{
		Username: "updated",
		Email:    "new@mail.com",
		Password: "pass", 
	}
	body, _ := json.Marshal(updatedData)

	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedUser models.User
	database.DB.First(&updatedUser, user.ID)
	assert.Equal(t, "updated", updatedUser.Username)
	assert.Equal(t, "new@mail.com", updatedUser.Email)
}


func TestDeleteUser(t *testing.T) {
	initUserTestDB()
	router := setupUserTestRouter()

	user := models.User{
		Username: "deletable",
		Email:    "delete@example.com",
		Password: "deletepass",
	}
	database.DB.Create(&user)

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var check models.User
	err := database.DB.First(&check, user.ID).Error
	assert.Error(t, err)
}
