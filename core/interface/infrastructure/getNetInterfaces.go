package interfaceInfra

import "net"

func GetNetInterfaces() ([]net.Interface, error) {
	return net.Interfaces()
}
