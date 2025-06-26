package entity_example_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
)

type mockBusiness struct{}

func (m *mockBusiness) Find() (*entity_example.EntityExample, error) {
	return nil, nil
}

func (m *mockBusiness) Create(c entity_example.EntityExample) (*entity_example.EntityExample, error) {
	return nil, nil
}

func (m *mockBusiness) Update(c entity_example.EntityExample) (*entity_example.EntityExample, error) {
	return nil, nil
}

func (m *mockBusiness) Delete(c entity_example.EntityExample) error {
	return nil
}

func TestEntityExampleController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockBusiness{}
	controller := entity_example.NewEntityExampleController(mock)

	t.Run("Find", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		controller.Find(ctx)
	})

	t.Run("Create", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		controller.Create(ctx)
	})

	t.Run("Update", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		controller.Update(ctx)
	})

	t.Run("Delete", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		controller.Delete(ctx)
	})
}
