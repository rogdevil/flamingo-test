package domain

type Product struct {
	ID    int
	Name  string
	Price float64
}

type ProductRepository interface {
	FindByID(id int) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
}
