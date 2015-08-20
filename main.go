package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello GitHub")

	fmt.Println(os.Getenv("PATH"))
}
