// controllers/protected_test.go

package controllers

import (
	"authapp/database"
	"authapp/models"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProfile(t *testing.T) {
	var profile models.User

	err := database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&models.User{})

	user := models.User{
		Email:    "jwt@email.com",
		Password: "secret",
		Name:     "Test User",
	}

	err = user.HashPassword(user.Password)
	assert.NoError(t, err)

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	request, err := http.NewRequest("GET", "/api/protected/profile", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	c.Set("email", "jwt@email.com")

	Profile(c)

	err = json.Unmarshal(w.Body.Bytes(), &profile)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

	log.Println(profile)

	assert.Equal(t, user.Email, profile.Email)
	assert.Equal(t, user.Name, profile.Name)
}

func TestProfileNotFound(t *testing.T) {
	var profile models.User

	err := database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&models.User{})

	request, err := http.NewRequest("GET", "/api/protected/profile", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	c.Set("email", "notfound@email.com")

	Profile(c)

	err = json.Unmarshal(w.Body.Bytes(), &profile)
	assert.NoError(t, err)

	assert.Equal(t, 404, w.Code)

	database.GlobalDB.Unscoped().Where("email = ?", "jwt@email.com").Delete(&models.User{})
}
