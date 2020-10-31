package main

import "fmt"

type IPAddr [4]byte

func (add IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", add[0], add[1], add[2], add[3])
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
