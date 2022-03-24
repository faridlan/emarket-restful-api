package web

type ProductResponse struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Price        int    `json:"price,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
	CategoryId   int    `json:"category_id,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
