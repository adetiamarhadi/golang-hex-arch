package main

import (
	"fmt"

	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/core/arithmetic"
)

func main() {
	
	arithmetic := arithmetic.NewAdapter()
	result, err := arithmetic.Add(1, 3)
	if err == nil {
		fmt.Println(result)
	}
}