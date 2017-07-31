package metrics

import (
	"github.com/rcrowley/go-metrics"
	"time"
	"os"
	"log"
)

func Metricsx() {
	m := metrics.NewMeter()
	metrics.Register("quux", m)
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	for s := 0; s < 30; s++ {
		for i := 0; i < 100; i++ {
			m.Mark(1)
		}
		time.Sleep(time.Second)
	}

}
