package barchart

import (
	"OCG_final_project_BE/api/model"
	"fmt"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type Test struct {
	Name  string
	Value float64
}

type Charter interface {
	Chart()
}

func CreateChartProduct(prt []model.ProductReport) (string, error) {
	values := []chart.Value{}
	for _, v := range prt {
		p := chart.Value{
			Value: v.Total,
			Label: v.Name,
		}
		values = append(values, p)
	}
	date := time.Now().Format("02-01-2006")

	graph := chart.BarChart{
		Title: "Top 10 best selling product   " + date,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		Height:   512,
		BarWidth: 50,
		Bars:     values,
	}
	path := fmt.Sprintf("./system/barchart/Product-" + date + ".png")
	f, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer f.Close()
	graph.Render(chart.PNG, f)
	return path, nil
}

func CreateChartOrder(prt model.OrderReport) (string, error) {
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
			{Value: float64(prt.TotalOrders), Label: "Total orders"},
			{Value: float64(prt.PaidOrders), Label: "Paid orders"},
			{Value: float64(prt.UnpaidOrders), Label: "Unpaid orders"},
		},
	}
	path := fmt.Sprintf("./system/barchart/OrderChart-" + date + ".png")
	f, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer f.Close()
	graph.Render(chart.PNG, f)
	return path, nil
}
