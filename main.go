package main

import (
	"fmt"
	"os"
)

const iteration = 3

func main() {
	fmt.Println("Iteration", iteration)
	fmt.Println("FLOX_VAR_TEST", os.Getenv("FLOX_VAR_TEST"))
}
