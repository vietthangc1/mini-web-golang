package main

type product struct {
	ID          string  `json:"id"`
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Number      int64   `json:"number"`
	Description string  `json:"description"`
	Cate1       string  `json:"cate1"`
	Cate2       string  `json:"cate2"`
	Color       string  `json:"color"`
	Size        string  `json:"size"`
}

