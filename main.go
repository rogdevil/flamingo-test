package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo-commerce-adapter-standalone/commercesearch"
	"flamingo.me/flamingo-commerce-adapter-standalone/csvindexing"
	"flamingo.me/flamingo-commerce/v3/cart"
	"flamingo.me/flamingo-commerce/v3/category"
	"flamingo.me/flamingo-commerce/v3/product"
	"flamingo.me/flamingo/v3"
	helloworld "github.com/rogdevil/flamingo-test/src/helloWorld"
)

func main() {
	flamingo.App(
		[]dingo.Module{
			new(helloworld.Module),
			// new(agoproduct.Module),
			new(product.Module),
			new(category.Module),
			new(cart.Module),
			new(csvindexing.ProductModule),
			new(commercesearch.Module),
			new(commercesearch.CategoryModule),
			new(commercesearch.SearchModule),
		},
	)
}
