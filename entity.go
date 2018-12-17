package translator

type ProductEntity struct {
	ID         int
	Name       string
	Price      int
	Cost       int
	CategoryID int
}

type CategoryEntity struct {
	ID   int
	Name string
}
