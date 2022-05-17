package main

import (
	"fmt"
	"math"
)

func main() {
	pi := math.NaN()
	fmt.Println(math.IsNaN(pi))
}
