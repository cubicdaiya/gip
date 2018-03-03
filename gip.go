package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var (
		ipStr   string
		cidrStr string
	)

	flag.StringVar(&ipStr, "ip", "", "IP address")
	flag.StringVar(&cidrStr, "cidr", "", "CIDR")
	flag.Parse()

	if ipStr == "" || cidrStr == "" {
		flag.Usage()
		os.Exit(1)
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		flag.Usage()
		os.Exit(1)
	}
	_, ipnet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		panic(err)
	}

	if ipnet.Contains(ip) {
		fmt.Printf("%s includes %s.", ipnet.String(), ip.String())
	} else {
		fmt.Printf("%s does not include %s.", ipnet.String(), ip.String())
	}
	fmt.Println()
}
