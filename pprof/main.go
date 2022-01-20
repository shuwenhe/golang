package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// http://127.0.0.1:6069/debug/pprof/
func main() {
	// 开启pprof，监听请求
	ip := "127.0.0.1:6069"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
