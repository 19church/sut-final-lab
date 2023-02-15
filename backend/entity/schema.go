package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name can not be blank"`
	Email      string `gorm:"uniqueIndex"`
	CustomerID string `gorm:"uniqueIndex" valid:"matches(^[LMH]\\d{7}$)"`
}
