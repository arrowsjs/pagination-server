package db

type Item struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	SubCategory   string  `json:"sub_category"`
	ContainerType string  `json:"container_type"`
	PricePerUnit  float64 `json:"price_per_unit"`
	Margin        float64 `json:"margin"`
}
