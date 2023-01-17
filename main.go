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

var (
	priceGuage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "degussa",
			Subsystem:   "listing",
			Name:        "euro",
			Help:        "The price of a Degussa item",
			ConstLabels: map[string]string{},
		},
		[]string{"name", "itemno", "price"},
	)
)

func recordMetrics() error {
	items, err := degussa.All()
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return errors.New("no items retrieved")
	}

	for _, item := range items {
		b := float64(item.BuyPrice) / 100
		s := float64(item.SellPrice) / 100

		priceGuage.WithLabelValues(item.Name, item.ItemNo, "buy").Set(b)
		priceGuage.WithLabelValues(item.Name, item.ItemNo, "sell").Set(s)
	}

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
