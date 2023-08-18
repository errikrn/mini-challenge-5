package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey"`
	VariantName string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	ProductID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Variant before create()")

	if len(v.VariantName) < 2 {
		err = errors.New("Variant name is too short")
	}

	return
}
