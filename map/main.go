package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Richard"] = 8
	m["Ritchie"] = 5
	m["ShuwenHe"] = 34
	fmt.Println("m = ", m)

	v1 := m["ShuwenHe"]
	fmt.Println("v1 = ", v1)

	fmt.Println("len = ", len(m))

	delete(m, "ShuwenHe")
	fmt.Println("m_del = ", m)

	v, ok := m["Ritchie"]
	fmt.Println("v  = ", v)
	fmt.Println("ok  = ", ok)

	family := map[string]int{"Sophie": 1, "Shuwen": 2}
	fmt.Println("family = ", family)
}
