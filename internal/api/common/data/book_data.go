package data

type BookData struct {
	ID     int32    `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Isbn   string   `json:"isbn"`
	Price  *float64 `json:"price,omitempty"`
}
