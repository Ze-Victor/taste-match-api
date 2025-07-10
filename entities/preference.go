package entities

type Preference struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"unique;not null"`
}
