package modelorder

import (
	"time"
)

type Order struct {
	OrderID      int       `json:"order_id" gorm:"primaryKey;autoIncrement"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"item" gorm:"constraint:OnDelete:CASCADE;"`
}

type Item struct {
	ItemId      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}

func (o *Order) SetFromRequest(orderRequest OrderRequest) {
	o.CustomerName = orderRequest.CustomerName

	for _, item := range orderRequest.Items {
		newItem := Item{}
		newItem.ItemCode = item.ItemCode
		newItem.Description = item.Description
		newItem.Quantity = item.Quantity

		o.Items = append(o.Items, newItem)
	}
}
