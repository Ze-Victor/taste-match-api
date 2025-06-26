package seed

import (
	"fmt"

	"gitlab.luizalabs.com/taste-match-api/api/entity_example"
	"gorm.io/gorm"
)

func CreadSeeds(db *gorm.DB) {
	createEntityExampleSeeds(db)
}

func createEntityExampleSeeds(db *gorm.DB) {

	records := []entity_example.EntityExample{
		{
			ID:   1,
			Name: "fulano",
		},
		{
			ID:   2,
			Name: "cavernicola",
		},
	}

	countInsertedRecords := 0
	for _, record := range records {
		if err := db.Create(&record).Error; err != nil {
			fmt.Printf("%v", err)
		} else {
			countInsertedRecords++
		}

	}

	fmt.Println("[Method: seed.Parametrizations()] Inserted records:", countInsertedRecords)

}
