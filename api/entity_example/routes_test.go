package entity_example_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
)

type mockEntityExampleController struct{}

func (m *mockEntityExampleController) Find(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"route": "find"})
}
func (m *mockEntityExampleController) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"route": "create"})
}
func (m *mockEntityExampleController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"route": "update"})
}
func (m *mockEntityExampleController) Delete(c *gin.Context) { c.Status(http.StatusNoContent) }

func TestSetupRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	controller := &entity_example.EntityExampleController{}
	entity_example.SetupRoutes(router.Group("/"), controller)

	req := httptest.NewRequest(http.MethodGet, "/address/search", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	req = httptest.NewRequest(http.MethodPost, "/address", bytes.NewBufferString(`{}`))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	req = httptest.NewRequest(http.MethodPatch, "/address/123", bytes.NewBufferString(`{}`))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	req = httptest.NewRequest(http.MethodDelete, "/address/123", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
