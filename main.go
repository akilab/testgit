package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
)

func main() {
	procs, _ := ps.Processes()
	for i, c := range procs {
		fmt.Println(i, c)
	}
	fmt.Println("Hello ブランチ")
}
