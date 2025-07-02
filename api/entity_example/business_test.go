package entity_example_test

import (
	"errors"
	"testing"

	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example"
)

type mockRepository struct {
	findFunc   func() (*entity_example.EntityExample, error)
	createFunc func(entity_example.EntityExample) (*entity_example.EntityExample, error)
	updateFunc func(entity_example.EntityExample) (*entity_example.EntityExample, error)
	deleteFunc func(entity_example.EntityExample) error
}

func (m *mockRepository) Find() (*entity_example.EntityExample, error) {
	return m.findFunc()
}

func (m *mockRepository) Create(c entity_example.EntityExample) (*entity_example.EntityExample, error) {
	return m.createFunc(c)
}

func (m *mockRepository) Update(c entity_example.EntityExample) (*entity_example.EntityExample, error) {
	return m.updateFunc(c)
}

func (m *mockRepository) Delete(c entity_example.EntityExample) error {
	return m.deleteFunc(c)
}

func TestEntityExampleBusinessImpl(t *testing.T) {
	mockRepo := &mockRepository{
		findFunc: func() (*entity_example.EntityExample, error) {
			return &entity_example.EntityExample{}, nil
		},
		createFunc: func(c entity_example.EntityExample) (*entity_example.EntityExample, error) {
			return &c, nil
		},
		updateFunc: func(c entity_example.EntityExample) (*entity_example.EntityExample, error) {
			return &c, nil
		},
		deleteFunc: func(c entity_example.EntityExample) error {
			return nil
		},
	}

	business := entity_example.NewEntityExampleBusinessImpl(mockRepo)

	if _, err := business.Find(); err != nil {
		t.Errorf("Find() error = %v", err)
	}

	entity := entity_example.EntityExample{}
	if _, err := business.Create(entity); err != nil {
		t.Errorf("Create() error = %v", err)
	}

	if _, err := business.Update(entity); err != nil {
		t.Errorf("Update() error = %v", err)
	}

	if err := business.Delete(entity); err != nil {
		t.Errorf("Delete() error = %v", err)
	}

	mockRepo.findFunc = func() (*entity_example.EntityExample, error) {
		return nil, errors.New("find error")
	}
}
