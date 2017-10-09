package main

import (
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	Calculate(1.1234, time.Now())
	time.Sleep(1 * time.Second)
	Calculate(1.1235, time.Now())
	time.Sleep(1 * time.Second)
	Calculate(1.1236, time.Now())
	time.Sleep(1 * time.Second)
	Calculate(1.1237, time.Now())

}
