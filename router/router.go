package router

import (
	"backend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/api/products", controller.GetProducts)
	r.HandleFunc("/api/products/{id:[0-9]+}", controller.GetOneProduct)

	r.HandleFunc("/api/pages/collections", controller.GetAllPages)
	r.HandleFunc("/api/pages/{id:[0-9]+}/collections", controller.GetOnePage)

	//dang tim cach lam
	r.HandleFunc("/api/collections/{id:[0-9]+}/categories", controller.GetOneCollection)

	r.HandleFunc("/api/categories/{id:[0-9]+}/products", controller.GetOneCategoryProducts)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9911", r))
}
