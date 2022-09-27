package params

import (
	"assignment2/pkg/models"
	"time"
)

type CreateOder struct {
	OrderID      int    `json:"order_id"`
	CustomerName string `json:"customerName"`
	Items        []CreateItem `json:"items"`
	OrderedAt    *time.Time `json:"orderedAt"`
}

type Orders struct {
	Order models.Order
	Items models.Item
}