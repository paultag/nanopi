package platform

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type LED string

var (
	BlueLED  LED = "nanopi:blue:status"
	GreenLED LED = "nanopi:green:pwr"
)

func ledPath(led LED) string {
	return fmt.Sprintf("/sys/class/leds/%s/brightness", led)
}

func GetLED(led LED) (bool, error) {
	fd, err := os.Open(ledPath(led))
	if err != nil {
		return false, err
	}
	defer fd.Close()
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		return false, err
	}

	return !(bytes.Compare(data, []byte{'0', '\n'}) == 0), nil
}

func SetLED(led LED, power bool) error {
	fd, err := WriteLED(led)
	if err != nil {
		return err
	}
	defer fd.Close()
	d := "0"
	if power {
		d = "1"
	}
	_, err = fd.Write([]byte(d))
	return err
}

func WriteLED(led LED) (*os.File, error) {
	return os.Create(ledPath(led))
}
