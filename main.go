package main

import (
	"log"

	"github.com/josedelrio85/test/prefix_sum"
)

func main() {
	log.Println("starting!")

	callPrefixSum()
}

func callPrefixSum() {
	log.Println("calling PrefixSum package")
	prefix_sum.Launch()
}
