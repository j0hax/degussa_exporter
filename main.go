package main

import (
	"log"
	"net/http"
	"time"

	"github.com/j0hax/degussa"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	kr := degussa.FilterTable(func(i degussa.Item) bool {
		return i.Name == "1 oz Krügerrand Goldmünze - Südafrika verschiedene Jahrgänge"
	})[0]

	b := float64(kr.BuyPrice) / 100
	s := float64(kr.SellPrice) / 100

	buyPriceGauge.Set(b)
	sellPriceGuage.Set(s)

	log.Println("Fetched new prices")
}

var (
	buyPriceGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "krugerrand_buy_price",
		Help: "The buy price of a 1 oz Krügerrand",
	})

	sellPriceGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "krugerrand_sell_price",
		Help: "The sell price of a 1 oz Krügerrand",
	})
)

func main() {
	go func() {
		for {
			recordMetrics()
			time.Sleep(5 * time.Minute)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":7979", nil)
}
