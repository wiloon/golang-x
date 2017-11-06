package main

import (
	"time"
	"log"
	"github.com/shopspring/decimal"
)

var bids []Price

func init() {

}

func Calculate(bid float64, timestamp time.Time) {
	log.Printf("calc, bid:%v, timestamp:%v", bid, timestamp)

	decimalBid := decimal.NewFromFloat(bid)

	bids = append(bids, Price{decimalBid, timestamp})

	for _, p := range bids {
		log.Println(p)
		a := decimalBid.Sub(p.bid)
		b := timestamp.UnixNano() - p.timestamp.UnixNano()

		log.Println("bid:", a)
		log.Println("t:", b/1e9)
	}
}

type Price struct {
	bid       decimal.Decimal
	timestamp time.Time
}
