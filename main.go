package main

import (
	"github.com/jghoman/gorestmath"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", gorestmath.DoSomeMath)
	if err := http.ListenAndServe("0.0.0.0:7552", nil); err != nil {
		log.Fatal("Couldn't start server", err)
	}
}
