package main

import (
	"fmt"

	dep "github.com/kevinburkemeter/vendor-repro-dependency"
	"github.com/kevinburkemeter/vendor-repro-dependency/subdep"
)

func main() {
	fmt.Println(dep.A)
	fmt.Println(subdep.B)
	//m := metrics.SimpleMetrics{}
	//fmt.Println(m)
}
