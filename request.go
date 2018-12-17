package translator

type CreateProductRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Cost       int    `json:"cost"`
	CategoryID int    `json:"category_id"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}
