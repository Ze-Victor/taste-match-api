package entity_example

import "time"

type EntityExample struct {
	ID        int        `gorm:"column:id; type:int; primaryKey"`
	Name      string     `gorm:"column:name; type:varchar(50); not null"`
	DeletedAt *time.Time `gorm:"column:deleted_at; type:datetime"`
}

func (b *EntityExample) TableName() string {
	return "entity_example"
}
