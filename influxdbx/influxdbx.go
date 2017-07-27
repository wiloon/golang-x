package main

import (
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"wiloon.com/golang-x/config"
)

func InfluxDbX() {
	influxdbAddress := config.GetString("influxdb.address")
	databaseName := config.GetString("influxdb.database.name")
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://" + influxdbAddress,


	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: databaseName,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	batchPoints.AddPoint(pt)

	// Write the batch
	if err := c.Write(batchPoints); err != nil {
		log.Fatal(err)
	}
}
