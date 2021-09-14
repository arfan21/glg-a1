package modelorder

type OrderRequest struct {
	CustomerName string        `json:"customer_name"`
	Items        []ItemRequest `json:"items"`
}

type ItemRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
