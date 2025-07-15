package main

import (
	"github.com/NoamFav/Zvezda/src/dashboard"
	"github.com/NoamFav/Zvezda/src/graph"
)

func main() {
	graph.Get_days()
	dashboard.Start()
}
