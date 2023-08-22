package products

import (
	"encoding/json"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/yoel0110/StockWave-API/internal/database"
)

type Products struct {
	Product_id          int8   `json:"productId"`
	Product_name        string `json:"productName"`
	Product_description string `json:"productDescription"`
	Product_category    string `json:"productCategory"`
}

type Response struct {
	ProductId int8
	Status    string
}

func CreateProduct(ctx *routing.Context) error {

	db := database.GetConnectDb()

	var product = new(Products)
	json.Unmarshal(ctx.Request.Body(), product)

	db.Exec("INSERT INTO products (product_name, product_description, product_category) VALUES (?, ?, ?)", product.Product_name, product.Product_description, product.Product_category)
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return ctx.WriteData("Ok")
}
