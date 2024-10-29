package orderbus

// Order representa o modelo de dados de uma ordem.
type Order struct {
	ID       int64   `json:"id"`
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
