package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func usage() {
	flag.Usage()
	os.Exit(1)
}

func printResult(ipnet *net.IPNet, ip net.IP) {
	if ipnet.Contains(ip) {
		fmt.Printf("%s includes %s.\n", ipnet.String(), ip.String())
	} else {
		fmt.Printf("%s does not include %s\n", ipnet.String(), ip.String())
	}
}

func main() {
	var (
		ipStr   string
		cidrStr string
	)

	flag.StringVar(&ipStr, "ip", "", "IP address")
	flag.StringVar(&cidrStr, "cidr", "", "CIDR")
	flag.Parse()

	if ipStr == "" || cidrStr == "" {
		usage()
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		usage()
	}
	_, ipnet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		log.Fatal(err)
	}

	printResult(ipnet, ip)
}
