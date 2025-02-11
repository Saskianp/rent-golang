package models 

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID            uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID    uint        `json:"customer_id"`
	Customer      Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
	CarID         uint        `json:"car_id"`
	Car           Car         `json:"car" gorm:"foreignKey:CarID"`
	StartRent     time.Time   `json:"start_rent"`
	EndRent       time.Time   `json:"end_rent"`
	TotalCost     float64     `json:"total_cost"` // Total cost dihitung otomatis
	Discount      float64     `json:"discount"`   // Diskon dari membership
	BookingTypeID uint        `json:"booking_type_id"`
	BookingType   BookingType `json:"booking_type" gorm:"foreignKey:BookingTypeID"`
	DriverID      *uint       `json:"driver_id" gorm:"null"` // Bisa null jika tanpa driver
	Driver        *Driver     `json:"driver" gorm:"foreignKey:DriverID"`
	DriverCost    float64     `json:"driver_cost"` // Biaya driver per hari
	Finished      bool        `json:"finished"`
}

// Function untuk menghitung TotalCost, Discount, dan DriverCost
func (b *Booking) CalculateTotalCost(db *gorm.DB) {
	var car Car
	var customer Customer
	var membership Membership
	var driver Driver

	// Ambil harga mobil
	db.First(&car, b.CarID)
	pricePerDay := car.DailyRent

	// Hitung durasi sewa
	duration := b.EndRent.Sub(b.StartRent).Hours() / 24
	if duration < 1 {
		duration = 1 // Minimal 1 hari
	}

	// Ambil informasi customer dan membership
	db.Preload("Membership").First(&customer, b.CustomerID)
	b.Customer = customer // Set customer agar Membership ikut terhubung

	// Jika customer memiliki membership, hitung diskon
	if customer.MembershipID != nil {
		db.First(&membership, *customer.MembershipID)
		b.Discount = (pricePerDay * duration) * (membership.Discount / 100)
	}

	// Ambil biaya driver jika ada
	if b.DriverID != nil {
		db.First(&driver, *b.DriverID)
		b.DriverCost = driver.DailyRent * duration
	}

	// Hitung total biaya (Total - Discount + DriverCost)
	b.TotalCost = (pricePerDay * duration) - b.Discount + b.DriverCost
}

// Migration untuk Booking
func MigrateBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
}
