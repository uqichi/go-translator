package translator

type AdminProductResponse struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Price    int            `json:"price"`
	Cost     int            `json:"cost"`
	Profit   int            `json:"profit"`
	Category *CategoryModel `json:"category,omitempty"`
}

type AppProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	//Cost     int            `json:"cost"`
	//Profit   int            `json:"profit"`
	Category *CategoryModel `json:"category,omitempty"`
}
