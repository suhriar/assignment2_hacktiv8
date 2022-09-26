package params

import "time"

type CreateOder struct {
	OrderID      int    `json:"order_id"`
	CustomerName string `json:"customerName"`
	Items        []CreateItem `json:"items"`
	OrderedAt    *time.Time `json:"orderedAt"`
}