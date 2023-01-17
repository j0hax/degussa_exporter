package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/j0hax/degussa"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// The name of the coin that is of tracking interest
const KrugerrandName = "1 oz Krügerrand Goldmünze - Südafrika verschiedene Jahrgänge"

func recordMetrics() error {
	items, err := degussa.FilterTable(func(i degussa.Item) bool {
		return i.Name == KrugerrandName
	})

	if err != nil {
		return err
	}

	if len(items) == 0 {
		return errors.New("no items retrieved")
	}

	kr := items[0]

	b := float64(kr.BuyPrice) / 100
	s := float64(kr.SellPrice) / 100

	buyPriceGauge.Set(b)
	sellPriceGuage.Set(s)

	log.Println("Fetched new prices")

	return nil
}

var (
	buyPriceGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "krugerrand_buy_euro",
		Help: "The current buy price of a 1 oz Krügerrand",
	})

	sellPriceGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "krugerrand_sell_euro",
		Help: "The current sell price of a 1 oz Krügerrand",
	})
)

func main() {
	go func() {
		for {
			err := recordMetrics()

			if err != nil {
				log.Panic(err)
				time.Sleep(time.Minute)
			} else {
				time.Sleep(5 * time.Minute)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":7979", nil)
}
