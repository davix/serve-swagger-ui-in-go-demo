package main

import (
	"main/simplepetstore"
	"net/http"

	api "github.com/go-openapi/runtime/middleware"
)

func main() {
	h, _ := simplepetstore.NewPetstore() // spec goes /swagger.json
	//uih := api.Redoc(api.RedocOpts{}, h) // redoc goes /docs
	uih := api.SwaggerUI(api.SwaggerUIOpts{}, h) // swagger ui goes /docs
	http.Handle("/", uih)
	http.ListenAndServe(":8344", nil)
}
