package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateenbagheri/mysmall-wallet/models"
	"github.com/mateenbagheri/mysmall-wallet/routes"
	"github.com/stretchr/testify/assert"
)

func TestAddTransaction1(t *testing.T) {
	r := routes.SetUpRouter()

	var transaction models.Transaction

	transaction.UserID = 4
	transaction.Amount = 500

	jsonValue, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAddTransaction2(t *testing.T) {
	r := routes.SetUpRouter()

	var transaction models.Transaction

	transaction.UserID = 5
	transaction.Amount = -500

	jsonValue, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAddTransaction3(t *testing.T) {
	r := routes.SetUpRouter()

	var transaction models.Transaction

	transaction.UserID = 4
	transaction.Amount = 0

	jsonValue, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
