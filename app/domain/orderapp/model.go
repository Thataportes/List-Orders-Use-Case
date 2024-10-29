package orderapp

type Order struct {
	ID       int64   `json:"id"`
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
