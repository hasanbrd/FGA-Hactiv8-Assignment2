package models

import "time"

type Order struct {
	OrderedAt    time.Time `json:"orderedAt"`
	OrderID      int       `gorm:"primaryKey" json:"-"`
	CustomerName string    `gorm:"type:varchar(20)" json:"customerName"`
	Items        []Item    `json:"items"`
}
