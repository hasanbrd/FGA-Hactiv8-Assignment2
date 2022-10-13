package models

type Item struct {
	ItemID      int    `gorm:"primaryKey" json:"itemID"`
	ItemCode    string `gorm:"type:varchar(20)" json:"itemCode"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Quantity    int    `gorm:"type:int" json:"quantity"`
	OrderID     int    `json:"-"`
}
