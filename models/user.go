package models

type User struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string
	Balance int64
}
