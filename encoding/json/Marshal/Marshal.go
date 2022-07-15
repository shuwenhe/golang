package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// m()
	m2()
}

func m() {
	b, _ := json.Marshal("gopher")
	fmt.Println(string(b))
}

func m2() {
	// sd := []string{"apple", "peach", "pear"}
	sd := []string{"apple"}
	fmt.Println("sd = ", sd)
	sb, _ := json.Marshal(sd)
	fmt.Println("sb = ", sb)
}
