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

func (r *ProductRepositoryImpl) FindAll(page, pageSize int) ([]*domain.Product, int, error) {
	fmt.Println("FindAll", page, pageSize)
	// Calculate offset
	offset := (page - 1) * pageSize

	// Get total count of products
	var totalCount int
	countQuery := "SELECT COUNT(*) FROM products"
	err := r.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated products
	query := "SELECT id, name, price FROM products LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
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
