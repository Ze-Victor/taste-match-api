package user

import "time"

type User struct {
	ID        int        `gorm:"column:id; type:int; primaryKey"`
	Name      string     `gorm:"column:name; type:varchar(50); not null"`
	Email     string     `gorm:"column:email; type:varchar(100); not null; unique"`
	Password  string     `gorm:"column:password; type:varchar(255); not null"`
	Gender    string     `gorm:"column:gender; type:varchar(1); not null"`
	BirthDate time.Time  `gorm:"column:birth_date; type:datetime; not null"`
	DeletedAt *time.Time `gorm:"column:deleted_at; type:datetime"`
}

func (b *User) TableName() string {
	return "users"
}
