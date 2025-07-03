package entity_example_test

import (
	"testing"
	"time"

	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example"
	"github.com/stretchr/testify/assert"
)

func TestEntityExample_TableName(t *testing.T) {
	now := time.Now()

	entity := &entity_example.EntityExample{
		ID:        1,
		Name:      "Test Entity",
		DeletedAt: &now,
	}

	assert.Equal(t, "entity_example", entity.TableName())
}
