package controller

import (
	"OCG_final_project_BE/api/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	od, err := model.Create(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	uj, err := json.Marshal(od)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "%s\n", uj)
	w.WriteHeader(http.StatusOK) // 200
}

func UpdateOrderPay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	od, err := model.UpdateAfterPay(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "Successfully update fulfill at  %v\n", od)
	w.WriteHeader(http.StatusOK) // 200
}

func OrderReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	st := vars["starttime"]
	et := vars["endtime"]
	od, err := model.OrderAnalysisClient(st, et)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	uj, err := json.Marshal(od)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "%s\n", uj)
	w.WriteHeader(http.StatusOK) // 200
}

func OrderManagement(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	st := vars["starttime"]
	et := vars["endtime"]
	od, err := model.OrderManagements(st, et)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	uj, err := json.Marshal(od)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
	}

	// ujt, err := json.Marshal(odt)
	// if err != nil {
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	fmt.Println(err)
	// }

	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "%s\n", uj)
	w.WriteHeader(http.StatusOK) // 200
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	err := model.UpdateOrderStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "Status: true")
	w.WriteHeader(http.StatusOK) // 200
}
