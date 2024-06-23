package data

type OrderItemData struct {
	ID            *int32   `json:"id,omitempty"`
	Book          BookData `json:"book"`
	Quantity      int32    `json:"quantity"`
	PurchasePrice *float64 `json:"purchase_price,omitempty"`
}
