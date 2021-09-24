package controller

import (
	"OCG_final_project_BE/api/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOneCategoryProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	limitStr := vars["limit"]
	limitNum, _ := strconv.Atoi(limitStr)

	offsetStr := vars["offset"]
	offsetNum, _ := strconv.Atoi(offsetStr)
	page, err := model.ProductsBasedCategories(id, limitNum, offsetNum)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(page)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "%s\n", uj)
	w.WriteHeader(http.StatusOK) // 200
}
