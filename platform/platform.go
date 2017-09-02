package platform

import (
	"encoding/hex"
	"io/ioutil"
	"os"
)

var (
	serialFile string = "/proc/device-tree/serial-number"
)

// Return the Serial in a byte array.
func Serial() ([]byte, error) {
	fd, err := os.Open(serialFile)
	if err != nil {
		return nil, err
	}

	d, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	// this platform has a trailing \0, so let's chop that dude off.
	d = d[:len(d)-1]

	dst := make([]byte, hex.DecodedLen(len(d)))
	_, err = hex.Decode(dst, d)
	if err != nil {
		return nil, err
	}
	return dst, nil
}
