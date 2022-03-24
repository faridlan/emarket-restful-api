package web

type ProductCreateRequest struct {
	Name       string `validate:"required,max=200,min=1" json:"name,omitempty"`
	Price      int    `json:"price,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
	CategoryId int    `json:"category_id,omitempty"`
}
