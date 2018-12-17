package translator

type ProductModel struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Price    int            `json:"price"`
	Cost     int            `json:"cost"`
	Profit   int            `json:"profit"`
	Category *CategoryModel `json:"category,omitempty"`
}

type CategoryModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
