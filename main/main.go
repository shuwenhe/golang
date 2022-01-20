package main

import "github.com/shuwenhe/shuwen-shop-admin/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail:", err.Error())
	}
}
