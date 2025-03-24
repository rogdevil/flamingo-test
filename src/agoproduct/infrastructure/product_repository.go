package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/rogdevil/ago/src/agoproduct/domain"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func (r *ProductRepositoryImpl) Inject() {
	r.db, _ = NewDB(&DBConfig{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     5432,
		Username: "agoAdmin",
		Password: "root@123",
		Database: "ago",
		SSLMode:  "disable",
	})
}

func NewProductRepository(db *sql.DB) *ProductRepositoryImpl {
	fmt.Println("Initializing ProductRepository with DB:", db)
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
	fmt.Println("Create", product.Name, product.Price, product.ID)
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id"
	fmt.Println(r.db)
	err := r.db.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		fmt.Println("Error creating product", err)
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
