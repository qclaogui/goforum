package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qclaogui/goforum/routes"
	"github.com/stretchr/testify/assert"
)

func TestInitRoutes(t *testing.T) {
	r := routes.InitRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
