package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range interfaces {
		fmt.Println("Interfaz:", iface.Name)
		fmt.Println("  MAC:", iface.HardwareAddr)

		addrs, err := iface.Addrs()
		if err != nil {
			log.Println("  Error obteniendo direcciones:", err)
			continue
		}

		for _, addr := range addrs {
			fmt.Println("  Dirección:", addr.String())
		}
		fmt.Println()
	}
}
