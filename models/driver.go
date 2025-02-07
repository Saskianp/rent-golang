package models

import "gorm.io/gorm"

type Driver struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string  `gorm:"type:varchar(100);not null" json:"name"`
	NIK         string `gorm:"type:varchar(16);unique;not null" json:"nik"`
	PhoneNumber string `gorm:"type:varchar(15);not null" json:"phoneNumber"`
	DailyRent float64 `gorm:"not null" json:"daily_rent"`
}

func MigrateDriver(db *gorm.DB) error {
	return db.AutoMigrate(&Driver{})
}
