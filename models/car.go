package models

import "gorm.io/gorm"

type Car struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string  `gorm:"type:varchar(100);not null" json:"name"`
	Stock     int     `gorm:"not null" json:"stock"`
	DailyRent float64 `gorm:"not null" json:"daily_rent"`
}

func MigrateCar(db *gorm.DB) error {
	return db.AutoMigrate(&Car{})
}
