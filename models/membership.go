package models

import "gorm.io/gorm"

type Membership struct {
	ID             uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	MembershipName string  `gorm:"type:varchar(100);not null" json:"membershipName"`
	Discount       float64 `gorm:"not null" json:"discount"`
}

func MigrateMembership(db *gorm.DB) error {
	return db.AutoMigrate(&Membership{})
}
