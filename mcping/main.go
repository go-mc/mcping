package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/go-mc/mcping"
	"github.com/mattn/go-colorable"
)

var protocol = flag.Int("protocol", 578, "The minecraft protocol version")
var output = colorable.NewColorableStdout()

func main() {
	flag.Parse()
	addr := flag.Arg(0)

	if _, _, err := net.SplitHostPort(addr); err != nil {
		if addrErr, ok := err.(*net.AddrError); ok && addrErr.Err == "missing port in address" {
			addr = net.JoinHostPort(lookupMC(addr))
		} else {
			exit(fmt.Errorf("split address error: %v", err))
		}
	}

	fmt.Fprintf(output, "MCPING (%s):\n", addr)
	status, delay, err := mcping.PingAndList(addr, *protocol)
	if err != nil {
		exit(err)
	}

	fmt.Fprintf(output,
		`server: %s
protocol: %d
description: 
%s
delay: %v
list: %d/%d
`,
		status.Version.Name,
		status.Version.Protocol,
		status.Description, delay,
		status.Players.Online,
		status.Players.Max)
	for _, v := range status.Players.Sample {
		fmt.Fprintf(output, "- [%s] %v\n", v.Name, v.ID)
	}
}

// written after read mojang's code
func lookupMC(addr string) (host, port string) {
	_, addrs, err := net.LookupSRV("minecraft", "tcp", addr)
	if err == nil && len(addrs) > 0 && addrs[0] != nil {
		return addrs[0].Target, strconv.Itoa(int(addrs[0].Port))
	}
	// lookup SRV error, return with default port
	return addr, "25565"
}

func exit(err error) {
	fmt.Fprintf(output, "error: %v\n", err)
	os.Exit(1)
}
