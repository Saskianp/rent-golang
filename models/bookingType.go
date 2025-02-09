package models

import "gorm.io/gorm"

type BookingType struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	BookingType string `gorm:"type:varchar(100);not null" json:"bookingType"`
	Description string    `gorm:"type:varchar(100);not null" json:"description"`
}

func MigrateBookingType(db *gorm.DB) error {
	return db.AutoMigrate(&BookingType{})
}
