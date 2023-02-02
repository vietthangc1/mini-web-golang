package main

import (
	"encoding/json"
)

type propertises struct {
	Color string `json:"color"`
	Brand string `json:"brand"`
	Size  string `json:"size"`
}
type product struct {
	ID          string      `json:"id"`
	SKU         string      `json:"sku"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Number      int64       `json:"number"`
	Description string      `json:"description"`
	Cate1       string      `json:"cate1"`
	Cate2       string      `json:"cate2"`
	Cate3       string      `json:"cate3"`
	Cate4       string      `json:"cate4"`
	Propertises propertises `json:"propertises"`
}

// Scan String to Struct for GET method
func parseJSONToModel(src interface{}, dest interface{}) error {
	var data []byte

	if b, ok := src.([]byte); ok {
		data = b
	} else if s, ok := src.(string); ok {
		data = []byte(s)
	} else if src == nil {
		return nil
	}

	return json.Unmarshal(data, dest)
}

func (p *propertises) Scan(src interface{}) error {
	return parseJSONToModel(src, p)
}

// Convert Propertise to String for POST method
func (p propertises) String() string {
	out, err := json.Marshal(p)
	if (err != nil) {
		panic(err)
	}
	return string(out)
}
