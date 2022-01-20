package main

import (
	"fmt"
)

func main() {
	BufferChannel()
}

func BufferChannel() {
	ch := make(chan string, 3) // 非阻塞：缓冲区通道为3
	fmt.Println("ch = ", ch)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	// ch <- "d" // fatal error: all goroutines are asleep - deadlock!
	count := len(ch)
	for i := 0; i < count; i++ { // 为什么for i:=0;i<len(ch);i++会少执行一次呢？
		fmt.Printf("<-ch%d = %s \n", i, <-ch)
	}
	// <-ch // fatal error: all goroutines are asleep - deadlock!
}
