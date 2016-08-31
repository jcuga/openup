package main

import (
	"flag"
	"fmt"
	"github.com/jcuga/go-upnp"
	"os"
)

func main() {
	displayExternalIpPtr := flag.Bool("ip", false, "Display external IP address and exit.")
	useUdpPtr := flag.Bool("udp", false, "Use UDP (instead of TCP) when opening/closing port forward.")
	doClosePtr := flag.Bool("close", false, "Close (as opposed to open) the given port.")
	portPtr := flag.Int("port", -1, "Port to open/close")
	flag.Parse()
	if *displayExternalIpPtr {
		// connect to router
		d, dErr := upnp.Discover()
		if dErr != nil {
			fmt.Printf("Error discovering router: %v\n", dErr)
			os.Exit(1)
		}
		// discover external IP
		ip, ipErr := d.ExternalIP()
		if ipErr != nil {
			fmt.Printf("Error fetching external IP address: %v\n", ipErr)
			os.Exit(1)
		}
		fmt.Println("Your external IP is:", ip)
		return
	}
	if *portPtr <= 0 {
		fmt.Printf("Invalid port number, must be > 0.  Got: %d\n", *portPtr)
		os.Exit(1)
	}
	d, dErr := upnp.Discover()
	if dErr != nil {
		fmt.Printf("Error discovering router: %v\n", dErr)
		os.Exit(1)
	}
	var proto string
	if *useUdpPtr {
		proto = "UDP"
	} else {
		proto = "TCP"
	}
	// Open/Close the given port/protocol
	if *doClosePtr {
		fmt.Printf("Closing port forward for %s port %d\n", proto, *portPtr)
		cErr := d.Clear(uint16(*portPtr), proto)
		if cErr != nil {
			fmt.Printf("Error closing port forward: %v\n", cErr)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Opening port forward for %s port %d\n", proto, *portPtr)
		fErr := d.Forward(uint16(*portPtr), "Requested by openup util.", proto)
		if fErr != nil {
			fmt.Printf("Error closing port forward: %v\n", fErr)
			os.Exit(1)
		}
	}
}
