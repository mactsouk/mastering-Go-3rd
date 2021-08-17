package main

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PORT is the TCP port number the server will listen to
var PORT = ":1234"

var n_goroutines = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_goroutines",
		Help:      "Number of goroutines"})

var n_memory = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_memory",
		Help:      "Memory usage"})

func main() {
	rand.Seed(time.Now().Unix())
	prometheus.MustRegister(n_goroutines)
	prometheus.MustRegister(n_memory)

	const nGo = "/sched/goroutines:goroutines"
	const nMem = "/memory/classes/heap/free:bytes"
	getMetric := make([]metrics.Sample, 2)
	getMetric[0].Name = nGo
	getMetric[1].Name = nMem

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			for i := 1; i < 4; i++ {
				go func() {
					_ = make([]int, 1000000)
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				}()
			}

			runtime.GC()
			metrics.Read(getMetric)
			goVal := getMetric[0].Value.Uint64()
			memVal := getMetric[1].Value.Uint64()
			time.Sleep(time.Duration(rand.Intn(15)) * time.Second)

			n_goroutines.Set(float64(goVal))
			n_memory.Set(float64(memVal))
		}
	}()

	log.Println("Listening to port", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}
