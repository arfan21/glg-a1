package modelorder

import "time"

type OrderRequest struct {
	OrderID      int           `json:"orderId,omitempty"`
	CustomerName string        `json:"customerName"`
	OrderedAt    time.Time     `json:"orderedAt"`
	Items        []ItemRequest `json:"items"`
}

type ItemRequest struct {
	ItemId      int    `json:"itemId,omitempty"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
