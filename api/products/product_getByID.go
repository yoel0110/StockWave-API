package products

import (
	"encoding/json"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/yoel0110/StockWave-API/internal/database"
)

type Product struct {
	Product_id int8 `json:"productId"`
}

func GetProductByID(ctx *routing.Context) error {
	db := database.GetConnectDb()

	var productId = new(Product)
	var product Products
	json.Unmarshal(ctx.Request.Body(), productId)
	log.Println(productId.Product_id)

	result := db.QueryRow("SELECT * FROM products WHERE product_id = ?", productId.Product_id)
	if err := result.Scan(&product.Product_id, &product.Product_name, &product.Product_description, &product.Product_description); err != nil {
		log.Println(err.Error())
		return ctx.WriteData("No se ha encontrado producto con la id especificada.")
	}

	jsonData, _ := json.Marshal(product)
	ctx.WriteData(jsonData)
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
