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
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string     `gorm:"type:varchar(100);not null" json:"name"`
	NIK          string     `gorm:"type:varchar(16);unique;not null" json:"nik"`
	PhoneNumber  string     `gorm:"type:varchar(15);not null" json:"phone_number"`
	MembershipID *uint      `json:"membership_id" gorm:"default:null"` // Bisa null jika customer tidak punya membership
	Membership   Membership `json:"membership" gorm:"foreignKey:MembershipID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func MigrateCust(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}
