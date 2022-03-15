package main

import "fmt"

func main() {
	c := 1
	a, b := ListAlbums(c)
	fmt.Println("a,b = ", a, b)
}

func ListAlbums(c int) (a, b int) {
	a = c
	b = a + 1
	return
}
