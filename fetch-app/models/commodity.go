package models

type Commodity struct {
	UUID           *string `json:"uuid"`
	Commodity      *string `json:"komoditas"`
	ProvincialArea *string `json:"area_provinsi"`
	CityArea       *string `json:"area_kota"`
	Size           *string `json:"size"`
	Price          *string `json:"price"`
	PriceUSD       *string `json:"price_usd,omitempty"`
	ParsedDate     *string `json:"tgl_parsed"`
	Timestampp     *string `json:"timestamp"`
}
