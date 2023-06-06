package main

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	SKU       string `json:"sku"`
	TimeStamp string `json:"time_stamp"`
}
