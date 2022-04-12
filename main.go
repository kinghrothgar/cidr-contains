package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var (
	helpFlag *bool
	helpHelp = "print help"

	quietFlag *bool
	quietHelp = "quiet; do not write anything to standard output"

	helpMSG = `Usage: %s <IP_CIDRs> <IP>

The first CIDR in list that containes given IP is returned with an exit code of 0
If there is no match, exit code 1 is returnned
If there is an error, exit code 2 is returned

Positional Arguments (Required):
- IP_CIDRs is a comma delineated list of CIDRs, eg. 100.0.0.0/16,192.168.1.1/16
- IP is the address to search the the CIDRs for

Flag Arguments (Optional):
-h %s
-q %s
`
)

func init() {
	filename := filepath.Base(os.Args[0])
	helpMSG = fmt.Sprintf(helpMSG, filename, helpHelp, quietHelp)
	helpFlag = flag.Bool("h", false, helpHelp)
	quietFlag = flag.Bool("q", false, quietHelp)
}

func main() {
	l := log.New(os.Stderr, "", 0)
	flag.Parse()
	if *helpFlag {
		l.Print(helpMSG)
		os.Exit(0)
	}
	if flag.NArg() != 2 {
		l.Println("ERROR: requires two positonal arguments: <IP CIDRs> <IP>")
		os.Exit(2)
	}
	ip := net.ParseIP(flag.Arg(1))
	if ip == nil {
		l.Printf("ERROR: '%s' is not a valid IP address\n", flag.Arg(1))
		os.Exit(2)
	}
	for _, ipCIDR := range strings.Split(flag.Arg(0), ",") {
		_, ipNet, err := net.ParseCIDR(ipCIDR)
		if err != nil {
			l.Printf("ERROR: %s is not a valid CIDR\n", ipCIDR)
			os.Exit(2)
		}

		if ipNet.Contains(ip) {
			if !*quietFlag {
				fmt.Println(ipNet.String())
			}
			os.Exit(0)
		}
	}
	os.Exit(1)
}
