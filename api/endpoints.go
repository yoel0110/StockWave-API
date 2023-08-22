package api

import (
	routing "github.com/qiangxue/fasthttp-routing"

	apiauth "github.com/yoel0110/StockWave-API/api/apiAuth"
	"github.com/yoel0110/StockWave-API/api/products"
)

func Route() *routing.Router {

	router := routing.New()

	api := router.Group("/api")
	api.Use(apiauth.Auth())

	apiProducts := api.Group("/products")
	apiProducts.Post("/createproducts", products.CreateProduct)
	apiProducts.Get("/getproducts", products.GetProducts)
	apiProducts.Get("/getproductById", products.GetProductByID)
	apiProducts.Get("/getproductByCategory", products.GetProductByCategory)
	return router
}
