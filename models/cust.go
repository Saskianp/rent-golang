// package models

// type Customer struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	NIK         string `json:"nik"`
// 	PhoneNumber string `json:"phone_number"`
// }

package models

import "gorm.io/gorm"

type Customer struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100);not null"`
	NIK         string `gorm:"type:varchar(16);unique;not null"`
	PhoneNumber string `gorm:"type:varchar(15);not null"`
}

func MigrateCust(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}
