package router

import (
	"OCG_final_project_BE/api/controller"
	"OCG_final_project_BE/api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/home", http.StatusSeeOther)

}
func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.WithLogging)
	r.HandleFunc("/", index)
	// Return all products
	// Return one products selected
	// Return list product with two params which are limit, omit product and category id selected
	// Return top 10 best selling products
	r.HandleFunc("/api/products", controller.GetProducts)
	r.HandleFunc("/api/products/{id:[0-9]+}", controller.GetOneProduct)
	r.HandleFunc("/api/products/{starttime}/{endtime}/bestsellings", controller.TopProduct)

	// Return  products based on categories.id
	r.Path("/api/categories/{id:[0-9]+}/products").
		Queries("limit", "{limit}").Queries("offset", "{offset}").
		HandlerFunc(controller.ShowProductBasedCategoryID)
	// Return products based on collection.id
	r.Path("/api/collections/{id:[0-9]+}/products").
		Queries("limit", "{limit}").Queries("offset", "{offset}").
		HandlerFunc(controller.ShowProductBasedCollectionID)

	// Return all page and its collections
	// Return a page and its collections
	r.HandleFunc("/api/pages/collections", controller.GetAllPages)
	r.HandleFunc("/api/pages/{id:[0-9]+}/collections", controller.GetOnePage)

	// Return a collection and its categories
	r.HandleFunc("/api/collections/{id:[0-9]+}/categories", controller.GetOneCollection)

	// Post a order with body
	// Return orde analysis such as total sales, paid order, unpaid order with prams start time and end time query

	r.HandleFunc("/api/orders", controller.CreateOrder).Methods("POST")
	r.HandleFunc("/api/orders/payment", controller.HandlePayment).Methods("POST")
	r.HandleFunc("/api/orders/{id:[0-9]+}/fulfill", controller.UpdateOrderPay)
	r.HandleFunc("/api/orders/{starttime}/{endtime}/analysis", controller.OrderReport)
	r.HandleFunc("/api/orders/{starttime}/{endtime}/managements", controller.OrderManagement)
	r.HandleFunc("/api/orders/{id:[0-9]+}/status", controller.UpdateOrderStatus)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9944", r))
}
