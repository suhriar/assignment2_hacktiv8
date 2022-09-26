package models

type Item struct {
	ID     int `gorm:"primaryKey" json:"item_id"`
	ItemCode    string    `gorm:"not null" json:"item_code"`
	Description string `gorm:"not null" json:"description"`
	Quantity    int `gorm:"not null" json:"quantity"`
	OrderId     int
}