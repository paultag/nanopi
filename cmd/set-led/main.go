package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
	"pault.ag/go/nanopi/platform"
)

func ohshit(err error) {
	if err != nil {
		panic(err)
	}
}

func Heartbeat(led platform.LED) {
	fd, err := platform.WriteLED(led)
	ohshit(err)
	defer fd.Close()
	for {
		fd.Write([]byte("1"))
		ohshit(fd.Sync())
		time.Sleep(time.Second / 8)

		fd.Write([]byte("0"))
		ohshit(fd.Sync())
		time.Sleep(time.Second / 8)

		fd.Write([]byte("1"))
		ohshit(fd.Sync())
		time.Sleep(time.Second / 8)

		fd.Write([]byte("0"))
		ohshit(fd.Sync())
		time.Sleep(time.Second * 1)
	}
}

func Action(c *cli.Context) error {
	blue := c.Bool("blue")
	green := c.Bool("green")
	heartbeat := c.Bool("heartbeat")
	state := c.Bool("state")

	if blue {
		ohshit(platform.SetLED(platform.BlueLED, state))
	}

	if green {
		ohshit(platform.SetLED(platform.GreenLED, state))
	}

	if heartbeat {
		Heartbeat(platform.BlueLED)
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
		cli.BoolFlag{Name: "heartbeat"},
	}

	app.Run(os.Args)

}
