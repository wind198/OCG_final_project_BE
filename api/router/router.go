package router

import (
	"backend/api/controller"
	"backend/api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.WithLogging)
	r.HandleFunc("/", index)
	// Return all products
	// Return one products selected
	// Return list product with two params which are limit product and category id selected
	r.HandleFunc("/api/products", controller.GetProducts)
	r.HandleFunc("/api/products/{id:[0-9]+}", controller.GetOneProduct)
	r.HandleFunc("/api/products/{limit:[0-9]+}/{offset:[0-9]+}/categories/{id:[0-9]+}", controller.GetOneCategoryProducts)

	// Return all page and its collections
	// Return a page and its collections
	r.HandleFunc("/api/pages/collections", controller.GetAllPages)
	r.HandleFunc("/api/pages/{id:[0-9]+}/collections", controller.GetOnePage)

	// Return a collection and its categories
	r.HandleFunc("/api/collections/{id:[0-9]+}/categories", controller.GetOneCollection)

	// Post a order with body
	r.HandleFunc("/api/orders", controller.CreateOrder).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9911", r))
}
