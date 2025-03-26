package agoproduct

import (
	"database/sql"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/rogdevil/ago/src/agoproduct/domain"
	"github.com/rogdevil/ago/src/agoproduct/infrastructure"
	"github.com/rogdevil/ago/src/agoproduct/interfaces"
	"github.com/rogdevil/ago/src/agoproduct/middleware"
)

type Module struct{}

func (m *Module) Configure(injector *dingo.Injector) {
	// Bind database connection
	injector.Bind(new(*sql.DB)).ToProvider(infrastructure.NewDB)

	// Bind repository
	injector.Bind(new(domain.ProductRepository)).To(infrastructure.ProductRepositoryImpl{})

	web.BindRoutes(injector, new(routes))
}

// routes struct holds the controller
type routes struct {
	productController *interfaces.ProductController
}

// Inject dependencies into the routes struct
func (r *routes) Inject(productController *interfaces.ProductController) *routes {
	r.productController = productController
	return r
}

// Routes method registers the routes
func (r *routes) Routes(registry *web.RouterRegistry) {
	// Register routes for CRUD operations

	corsHandler := middleware.CorsHandler{}
	// Get all products (paginated)
	registry.MustRoute("/products/paginated", "product.getAllPaginated")
	registry.HandleOptions("product.getAllPaginated", web.WrapHTTPHandler(corsHandler.PreflightHandler()))
	registry.HandleGet("product.getAllPaginated", r.productController.GetProductsPaginated)

	// GET /products/:id - Get a product by ID
	registry.MustRoute("/products/:id", "product.get")
	registry.HandleGet("product.get", r.productController.GetProduct)

	// POST /products - Create a new product
	registry.MustRoute("/products", "product.create")
	registry.HandlePost("product.create", r.productController.CreateProduct)

	// PUT /products/:id - Update a product by ID
	registry.MustRoute("/products/:id", "product.update")
	registry.HandlePut("product.update", r.productController.UpdateProduct)

	// DELETE /products/:id - Delete a product by ID
	registry.MustRoute("/products/:id", "product.delete")
	registry.HandleDelete("product.delete", r.productController.DeleteProduct)
}
