package models

import (
	"encoding/json"
)

type Propertises struct {
	Color string `json:"color"`
	Brand string `json:"brand"`
	Size  string `json:"size"`
}
type Product struct {
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
	Propertises Propertises `json:"propertises"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Scan String to Struct for GET method
func ParseJSONToModel(src interface{}, dest interface{}) error {
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

func (p *Propertises) Scan(src interface{}) error {
	return ParseJSONToModel(src, p)
}

// Convert Propertise to String for POST method
func (p Propertises) String() string {
	out, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}
