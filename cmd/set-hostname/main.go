package main

import (
	"fmt"
	"syscall"

	"pault.ag/go/nanopi/platform"
)

func main() {
	id, err := platform.Serial()
	if err != nil {
		panic(err)
	}
	if err := syscall.Sethostname([]byte(fmt.Sprintf("nanopi-%x", id))); err != nil {
		panic(err)
	}
}
