package main

import (
	"flag"

	"github.com/painhardcore/WebHasher/pkg/hashfetcher"
)

const (
	parallelDefault = 10
)

func main() {
	parallelFlag := flag.Int("parallel", parallelDefault, "number of parallel requests")
	flag.Parse()
	hashfetcher.Run(*parallelFlag, flag.Args())
}
