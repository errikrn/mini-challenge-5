package models

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique;type:varchar(191)"`
	Variants  []Variant `gorm: "foreignKey:ProductID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
