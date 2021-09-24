package barchart

import (
	"OCG_final_project_BE/api/model"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

func CreateChartProduct(product []model.ProductReport) (string, error) {
	charts := make([]chart.Value, 0)
	fmt.Println(charts)
	for _, v := range product {
		p := chart.Value{
			Label: v.Name,
			Value: v.Total,
		}
		charts = append(charts, p)
	}

	date := time.Now().Format("02-01-2006")
	graph := chart.BarChart{
		Title: "Top 10 best selling product " + date,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		Height:   512,
		BarWidth: 50,
		Bars:     charts,
	}
	path := fmt.Sprintf("./system/barchart/product" + date + ".png")
	f, err := os.Create(path)
	if err != nil {
		return path, errors.New("can not create file with path %v" + path)
	}
	graph.Render(chart.PNG, f)
	defer f.Close()
	return path, nil
}

func CreateChartOrder(order model.OrderReport) (string, error) {
	date := time.Now().Format("02-01-2006")

	graph := chart.BarChart{
		Title: "Order Bar Chart :" + date,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		Height:   512,
		BarWidth: 50,
		Bars: []chart.Value{
			{Value: float64(order.TotalOrders), Label: "Total orders"},
			{Value: float64(order.PaidOrders), Label: "Paid orders"},
			{Value: float64(order.UnpaidOrders), Label: "Unpaid orders"},
		},
	}
	path := fmt.Sprintf("./system/barchart/order" + date + ".png")
	f, err := os.Create(path)
	if err != nil {
		return path, errors.New("can not create file with path %v" + path)
	}
	graph.Render(chart.PNG, f)
	defer f.Close()
	return path, nil
}
