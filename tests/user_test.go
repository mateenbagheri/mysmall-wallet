package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateenbagheri/mysmall-wallet/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetUserBalance1(t *testing.T) {
	mockResponse := `
		{
			"balance": -1000
		}
	`

	r := routes.SetUpRouter()

	req, _ := http.NewRequest("GET", "/user/1/balance", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserBalance2(t *testing.T) {
	mockResponse := `
		{
			"balance": 0
		}
	`

	r := routes.SetUpRouter()

	req, _ := http.NewRequest("GET", "/user/2/balance", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserBalance3(t *testing.T) {
	mockResponse := `
	{
		"balance": -1000
	}
	`

	r := routes.SetUpRouter()

	req, _ := http.NewRequest("GET", "/user/1/balance", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
