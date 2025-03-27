package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rogdevil/flamingo-test/src/agoproduct/domain"

	"flamingo.me/flamingo/v3/framework/web"
)

type ProductController struct {
	repo      domain.ProductRepository
	responder *web.Responder // Inject the responder
}

func (c *ProductController) Inject(repo domain.ProductRepository, responder *web.Responder) {
	responder.TODO().Header.Set("Access-Control-Allow-Origin", "*")
	responder.TODO().Header.Set("Access-Control-Allow-Credentials", "true")
	responder.TODO().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
	responder.TODO().Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	c.repo = repo
	c.responder = responder
}

// GetProduct handles GET /products/:id
func (c *ProductController) GetProduct(ctx context.Context, req *web.Request) web.Result {
	idStr := req.Params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 400,
			"body":   "Invalid product ID",
		})
	}

	product, err := c.repo.FindByID(id)
	if err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 404,
			"body":   "Product not found",
		})
	}

	return c.responder.Data(product)
}

// GetProductsPaginated handles GET /products/paginated?page=1&pageSize=10
func (c *ProductController) GetProductsPaginated(ctx context.Context, req *web.Request) web.Result {
	page, err := strconv.Atoi(req.Request().URL.Query().Get("page"))
	fmt.Println("page", page)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(req.Request().URL.Query().Get("pageSize"))
	fmt.Println("pageSize", pageSize)
	if err != nil || pageSize < 1 {
		pageSize = 10 // default page size
	}

	products, totalCount, err := c.repo.FindAll(page, pageSize)
	if err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 500,
			"body":   "Failed to fetch products",
		})
	}

	result := c.responder.Data(map[string]interface{}{
		"products":   products,
		"totalCount": totalCount,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": (totalCount + pageSize - 1) / pageSize,
	})
	result.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	result.Header.Set("Access-Control-Allow-Headers", "*")
	result.Header.Set("Access-Control-Allow-Origin", "*")

	fmt.Println("result", result)

	return result
}

// CreateProduct handles POST /products
func (c *ProductController) CreateProduct(ctx context.Context, req *web.Request) web.Result {
	fmt.Println("CreateProduct")
	var product domain.Product
	if err := json.NewDecoder(req.Request().Body).Decode(&product); err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 400,
			"body":   "Invalid request body",
		})
	}

	fmt.Println("product", product)

	if err := c.repo.Create(&product); err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 500,
			"body":   "Failed to create product",
		})
	}

	return c.responder.Data(product)
}

// UpdateProduct handles PUT /products/:id
func (c *ProductController) UpdateProduct(ctx context.Context, req *web.Request) web.Result {
	idStr := req.Params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 400,
			"body":   "Invalid product ID",
		})
	}
	var product domain.Product
	if err := json.NewDecoder(req.Request().Body).Decode(&product); err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 400,
			"body":   "Invalid request body",
		})
	}

	product.ID = id
	if err := c.repo.Update(&product); err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 500,
			"body":   "Failed to update product",
		})
	}

	return c.responder.Data(product)
}

// DeleteProduct handles DELETE /products/:id
func (c *ProductController) DeleteProduct(ctx context.Context, req *web.Request) web.Result {
	idStr := req.Params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 400,
			"body":   "Invalid product ID",
		})
	}

	if err := c.repo.Delete(id); err != nil {
		return c.responder.Data(map[string]interface{}{
			"status": 500,
			"body":   "Failed to delete product",
		})
	}

	return c.responder.Data(map[string]interface{}{
		"status": 204,
		"body":   "Product deleted successfully",
	})
}

func (c *ProductController) Options(ctx context.Context, req *web.Request) web.Result {
	result := c.responder.Data(map[string]interface{}{})

	result.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	result.Header.Set("Access-Control-Allow-Headers", "*")
	result.Header.Set("Access-Control-Allow-Origin", "*")

	return result
}
