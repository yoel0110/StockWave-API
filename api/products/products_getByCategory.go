package products

import (
	"encoding/json"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/yoel0110/StockWave-API/internal/database"
)

func GetProductByCategory(ctx *routing.Context) error {
	db := database.GetConnectDb()
	var category = new(Products)
	var productsByCategory []Products

	json.Unmarshal(ctx.Request.Body(), category)

	results, _ := db.Query("SELECT * FROM products WHERE product_category = ?", category.Product_category)
	for results.Next() {
		var product Products
		if err := results.Scan(&product.Product_id, &product.Product_name, &product.Product_description, &product.Product_description); err != nil {
			log.Println(err.Error())
		}
		productsByCategory = append(productsByCategory, product)
	}

	jsonData, _ := json.Marshal(productsByCategory)
	ctx.SetStatusCode(fasthttp.StatusAccepted)
	return ctx.WriteData(jsonData)
}
