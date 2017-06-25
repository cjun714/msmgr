package nic

import (
	"errors"
	"net"
	"os"
	"util/log"
)

// GetIPByNicName get specific NIC's IP address(IPv4)
func GetIPByNicName(nicName string) (string, error) {
	fs, _ := net.Interfaces()
	for _, f := range fs {
		log.D(f.Name)
	}

	nss, _ := net.InterfaceByName(nicName)
	dd, _ := nss.Addrs()
	for _, ad := range dd {
		if ipnet, ok := ad.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.H("nic IP:", ipnet.IP.String())
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("no IP address can be found")
}

// GetHostName is used to return host's name
func GetHostName() string {
	host, e := os.Hostname()
	if e != nil {
		log.E(e)
		return "localhost"
	}
	log.H("host name: ", host)
	return host
}
