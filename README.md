# degussa_exporter
Prometheus Exporter for Degussa Goldhandel prices

## Purpose
This exporter was specifically written to learn Prometheus' `client_golang` library, especially to serve as a testing ground for more advanced topics such as using [GuageVec](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/promauto#NewGaugeVec)s and [histograms](https://prometheus.io/docs/practices/histograms/).

## Exported Metrics
Metrics are available on port `7979` (the atomic number of gold lol) under the standard URL `/metrics`.

The program exports all listed item prices on Degussa's [price page](https://www.degussa-goldhandel.de/preise/preisliste/). Prices are scraped at the same rate that Degussa publishes them, ca. 5 minutes.

### Example
```
# HELP degussa_listing_euro The price of a Degussa item
# TYPE degussa_listing_euro gauge
degussa_listing_euro{itemno="100012/01",name="1 g Degussa Goldbarren (geprägt)",price="buy"} 56
degussa_listing_euro{itemno="100012/01",name="1 g Degussa Goldbarren (geprägt)",price="sell"} 72.4
degussa_listing_euro{itemno="100026/01",name="2,5 g Degussa Goldbarren (geprägt)",price="buy"} 140.25
degussa_listing_euro{itemno="100026/01",name="2,5 g Degussa Goldbarren (geprägt)",price="sell"} 166.4
degussa_listing_euro{itemno="100052/01",name="5 g Degussa Goldbarren (geprägt)",price="buy"} 280.5
degussa_listing_euro{itemno="100052/01",name="5 g Degussa Goldbarren (geprägt)",price="sell"} 316.5
```
