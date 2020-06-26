package main

import (
	"net/http"

	"go-commerce/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3001", nil)
}
