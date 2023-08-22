package products

import (
	"encoding/json"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/yoel0110/StockWave-API/internal/database"
)

func GetProducts(ctx *routing.Context) error {
	db := database.GetConnectDb()
	var products []Products

	results, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err.Error())
	}

	defer results.Close()

	for results.Next() {
		var product Products
		if err := results.Scan(&product.Product_id, &product.Product_name, &product.Product_description, &product.Product_category); err != nil {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			return ctx.WriteData("No se encontraron productos en el inventario.\n")
		}
		products = append(products, product)
	}

	jsonData, _ := json.Marshal(products)
	ctx.SetStatusCode(fasthttp.StatusOK)
	return ctx.WriteData(jsonData)
}
