package entity_example_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
)

func TestEntityExampleRepositoryImpl(t *testing.T) {
	repo := entity_example.NewEntityExampleRepository()

	t.Run("Find", func(t *testing.T) {
		entity, err := repo.Find()
		assert.Nil(t, entity)
		assert.Nil(t, err)
	})

	t.Run("Create", func(t *testing.T) {
		entity, err := repo.Create(entity_example.EntityExample{})
		assert.Nil(t, entity)
		assert.Nil(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		entity, err := repo.Update(entity_example.EntityExample{})
		assert.Nil(t, entity)
		assert.Nil(t, err)
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.Delete(entity_example.EntityExample{})
		assert.Nil(t, err)
	})
}
