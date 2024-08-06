package main

import (
	"net/http"

	"github.com/marin/api/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
