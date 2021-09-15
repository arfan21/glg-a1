package modelorder

import "time"

type OrderResponse struct {
	OrderID      int            `json:"order_id"`
	CustomerName string         `json:"customer_name"`
	OrderedAt    time.Time      `json:"ordered_at"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemId      int    `json:"item_id,omitempty"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (o *OrderResponse) Set(order Order) {
	o.OrderID = order.OrderID
	o.CustomerName = order.CustomerName
	o.OrderedAt = order.OrderedAt

	for _, item := range order.Items {
		newItem := ItemResponse{}
		newItem.ItemId = item.ItemId
		newItem.ItemCode = item.ItemCode
		newItem.Description = item.Description
		newItem.Quantity = item.Quantity

		o.Items = append(o.Items, newItem)
	}
}
