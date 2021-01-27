package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func request(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRoot(t *testing.T) {
	body := gin.H{
		"location": "root",
	}
	method := "GET"
	testEndpoint(t, "/", body, method)
}

func TestAdd(t *testing.T) {
	body := gin.H{
		"location": "add",
	}
	method := "GET"
	testEndpoint(t, "/add", body, method)
}

func TestRemove(t *testing.T) {
	body := gin.H{
		"location": "remove",
	}
	method := "DELETE"
	testEndpoint(t, "/remove", body, method)
}
func TestList(t *testing.T) {
	body := gin.H{
		"location": "list",
	}
	method := "GET"
	testEndpoint(t, "/list", body, method)
}

func testEndpoint(t *testing.T, path string, body gin.H, method string) {
	r := router()
	w := request(r, method, path)
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["location"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["location"], value)
}
