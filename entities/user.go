package entities

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	BirthDate time.Time `gorm:"type:date;not null"`
	CreatedAt time.Time

	Preferences []Preference `gorm:"many2many:user_preferences;"`
}
