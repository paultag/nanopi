package main

import (
	"log"
	"net"
	"os"

	"crypto/sha256"

	"pault.ag/go/debian/control"
	"pault.ag/go/macchanger"
	"pault.ag/go/nanopi/platform"
)

type Config struct {
	control.Paragraph

	Interface string
}

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

func LoadConfig() Config {
	config := Config{Interface: "eth0"}

	fd, err := os.Open("/etc/set-emac")
	if err != nil {
		return config
	}

	if err := control.Unmarshal(&config, fd); err != nil {
		return config
	}
	return config
}

func main() {
	log.Printf("Loading set-emac script\n")
	config := LoadConfig()
	log.Printf("Looking up interface %s\n", config.Interface)
	iface, err := net.InterfaceByName(config.Interface)
	if err != nil {
		panic(err)
	}
	log.Printf("Generating a stable MAC from the device Serial\n")
	mac, err := GenerateStableMAC()
	if err != nil {
		panic(err)
	}
	log.Printf("Changing the MAC Address to the %s\n", mac)
	if err := macchanger.ChangeHardwareAddr(*iface, *mac); err != nil {
		panic(err)
	}
	log.Printf("MAC address set to %s on iface %s\n", mac, iface)
}
