// controllers/public_test.go

package controllers

import (
	"authapp/database"
	"authapp/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	var actualResult models.User

	user := models.User{
		Name:     "Test User",
		Email:    "jwt@email.com",
		Password: "secret",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/signup", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&models.User{})

	Signup(c)

	assert.Equal(t, 200, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &actualResult)
	assert.NoError(t, err)

	assert.Equal(t, user.Name, actualResult.Name)
	assert.Equal(t, user.Email, actualResult.Email)
}

func TestSignUpInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/signup", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Signup(c)

	assert.Equal(t, 400, w.Code)
}

func TestLogin(t *testing.T) {
	user := LoginPayload{
		Email:    "jwt@email.com",
		Password: "secret",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&models.User{})

	Login(c)

	assert.Equal(t, 200, w.Code)

}

func TestLoginInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Login(c)

	assert.Equal(t, 400, w.Code)
}

func TestLoginInvalidCredentials(t *testing.T) {
	user := LoginPayload{
		Email:    "jwt@email.com",
		Password: "invalid",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&models.User{})

	Login(c)

	assert.Equal(t, 401, w.Code)

	database.GlobalDB.Unscoped().Where("email = ?", user.Email).Delete(&models.User{})
}
