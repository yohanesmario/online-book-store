package data

type OrderData struct {
	ID     *int32          `json:"id,omitempty"`
	UserID int32           `json:"user_id"`
	Items  []OrderItemData `json:"items"`
}
