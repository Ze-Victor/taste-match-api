package migration

import (
	"github.com/Ze-Victor/taste-match-api/taste-match-api/api/entity_example"
	"github.com/Ze-Victor/taste-match-api/taste-match-api/internal/config"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	if config.Get().Env == "test" {
		DropTables(db)
	}

	db.AutoMigrate(&entity_example.EntityExample{})

	if db.Error != nil {
		panic(db.Error)
	}
}

// DropTables will drop all database tables to recreate it again.
// It should be used just on test environment
func DropTables(db *gorm.DB) {

	db.AutoMigrate(&entity_example.EntityExample{})

}
