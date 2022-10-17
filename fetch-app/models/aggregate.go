package models

type Aggregate struct {
	ProvincialArea string    `json:"area_provinsi"`
	Year           int       `json:"year"`
	Week           int       `json:"week"`
	MinPrice       float64   `json:"min_price"`
	MaxPrice       float64   `json:"max_price"`
	MeanPrice      float64   `json:"mean_price"`
	MedianPrice    float64   `json:"median_price"`
	TotalPrice     float64   `json:"total_price"`
	ListSize       []float64 `json:"list_price"`
	MinSize        float64   `json:"min_size"`
	MaxSize        float64   `json:"max_size"`
	MeanSize       float64   `json:"mean_size"`
	MedianSize     float64   `json:"median_size"`
	TotalSize      float64   `json:"total_size"`
	ListPrice      []float64 `json:"list_size"`
}
