package main

import (
	"net"

	"crypto/sha256"

	"pault.ag/go/macchanger"
	"pault.ag/go/nanopi/platform"
)

func GenerateStableMAC() (*net.HardwareAddr, error) {
	id, err := platform.Serial()
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(id)
	hashId := h.Sum(nil)

	mac := net.HardwareAddr(append([]byte{0x0E}, hashId[0:5]...))
	return &mac, nil
}

func main() {
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		panic(err)
	}
	mac, err := GenerateStableMAC()
	if err != nil {
		panic(err)
	}
	if err := macchanger.ChangeHardwareAddr(*iface, *mac); err != nil {
		panic(err)
	}
}
