package models

import (
	"time"
)

type Order struct {
	ID      int       `gorm:"primaryKey" json:"order_id"`
	CustomerName string     `gorm:"not null" json:"customer_name" form:"customer_name" valid:"required~Your customer name is required"`
	OrderedAt    *time.Time `json:"ordered_at,omitempty"`
	Item         []Item     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}