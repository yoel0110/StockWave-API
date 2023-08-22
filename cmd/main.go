package main

import (
	"log"
	"os"

	"github.com/valyala/fasthttp"
	"github.com/yoel0110/StockWave-API/api"
	"github.com/yoel0110/StockWave-API/internal/database"
)

func main() {

	errMkdir := os.Mkdir("./logs", 0570)
	if errMkdir != nil && os.IsExist(errMkdir) {
		log.Println(errMkdir)
	}

	f, err := os.OpenFile("./logs/process.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte("StockWave API Logs\nRegistro de errores de procesos.\n"))

	database.DbConnect()
	log.Fatal(fasthttp.ListenAndServe(":8080", api.Route().HandleRequest))
}
