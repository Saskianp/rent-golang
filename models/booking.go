package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID uint      `json:"customer_id"`    
	Customer   Customer  `json:"customer" gorm:"foreignKey:CustomerID"` // Tambahkan relasi
	CarID      uint      `json:"car_id"`
	Car        Car       `json:"car" gorm:"foreignKey:CarID"` // Tambahkan relasi
	StartRent  time.Time `json:"start_rent"`
	EndRent    time.Time `json:"end_rent"`
	TotalCost  float64   `json:"total_cost"` // Total cost yang dihitung otomatis
	Finished   bool      `json:"finished"`
}

func MigrateBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
}
