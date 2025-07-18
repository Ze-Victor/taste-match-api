package preferences

type User struct {
	ID          int    `gorm:"column:id; type:int; primaryKey"`
	Description string `gorm:"column:name; type:varchar(50); not null"`
}

func (b *User) TableName() string {
	return "users"
}
