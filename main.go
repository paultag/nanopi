package main

import (
	"fmt"

	// "pault.ag/go/macchanger"
	"pault.ag/go/nanopi/platform"
)

func main() {
	id, err := platform.Serial()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", id)
}
