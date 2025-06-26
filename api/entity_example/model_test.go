package entity_example_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
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
