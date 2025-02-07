package models

import (
	"gorm.io/gorm"
)

type DriverIncentive struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingID  uint      `json:"booking_id"`    
	Booking    Booking   `json:"booking" gorm:"foreignKey:BookingID"` // Tambahkan relasi
	Incentive  float64   `json:"total_cost"` // Total cost yang dihitung otomatis
}

func MigrateDriverIncentive(db *gorm.DB) {
	db.AutoMigrate(&DriverIncentive{})
}
