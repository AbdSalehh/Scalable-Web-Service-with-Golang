package models

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"orderId"`
	CustomerName string `gorm:"not null" json:"customerName"`
	OrderedAt    string `gorm:"not null" json:"orderedAt"`
	Items        []Item `gorm:"foreignKey:OrderID" json:"items"`
}
