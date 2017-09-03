package main

import (
	"fmt"
	"time"

	"pault.ag/go/nanopi/platform"
)

func ohshit(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ohshit(platform.SetLED(platform.BlueLED, false))
	state, err := platform.GetLED(platform.BlueLED)
	ohshit(err)
	fmt.Printf("LED: %s\n", state)

	time.Sleep(time.Second * 1)

	ohshit(platform.SetLED(platform.BlueLED, true))
	state, err = platform.GetLED(platform.BlueLED)
	ohshit(err)
	fmt.Printf("LED: %s\n", state)

	time.Sleep(time.Second * 1)

	ohshit(platform.SetLED(platform.BlueLED, false))
	state, err = platform.GetLED(platform.BlueLED)
	ohshit(err)
	fmt.Printf("LED: %s\n", state)
}
