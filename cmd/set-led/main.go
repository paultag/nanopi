package main

import (
	"os"

	"github.com/urfave/cli"
	"pault.ag/go/nanopi/platform"
)

func ohshit(err error) {
	if err != nil {
		panic(err)
	}
}

func Action(c *cli.Context) error {
	blue := c.Bool("blue")
	green := c.Bool("green")
	state := c.Bool("state")

	if blue {
		ohshit(platform.SetLED(platform.BlueLED, state))
	}

	if green {
		ohshit(platform.SetLED(platform.GreenLED, state))
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "set-led"
	app.Usage = "set the LED on a NanoPi"
	app.Action = Action

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "state"},
		cli.BoolFlag{Name: "blue"},
		cli.BoolFlag{Name: "green"},
	}

	app.Run(os.Args)

}
