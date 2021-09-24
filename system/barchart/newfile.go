package barchart

import (
	"OCG_final_project_BE/api/model"
	"testing"
)

func Test_CreateChartProduct(t *testing.T) {
	productReport := []model.ProductReport{
		{ID: 1, Name: "a", Total: 1},
		{ID: 2, Name: "b", Total: 2},
		{ID: 3, Name: "c", Total: 3},
	}
	product, err := CreateChartProduct(productReport)
	if err != nil {
		t.Error(err)
	}
	t.Log(product)
}
