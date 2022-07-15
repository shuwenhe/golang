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
	sd := []string{"apple", "peach", "pear"}
	// sd := []string{"apple"}
	fmt.Println("sd = ", sd) // sd =  [apple peach pear]
	sb, _ := json.Marshal(sd)
	fmt.Println("sb = ", sb) // sb =  [91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93]
}
