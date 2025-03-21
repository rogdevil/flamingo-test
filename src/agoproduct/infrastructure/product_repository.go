package infrastructure

import (
	"database/sql"

	"github.com/rogdevil/ago/src/agoproduct/domain"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) FindByID(id int) (*domain.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = $1"
	row := r.db.QueryRow(query, id)

	product := &domain.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) Create(product *domain.Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) Update(product *domain.Product) error {
	query := "UPDATE products SET name = $1, price = $2 WHERE id = $3"
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	return err
}

func (r *ProductRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
