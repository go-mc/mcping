package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/go-mc/mcping"
)

var protocol = flag.Int("protocol", 578, "The minecraft protocol version")

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

	fmt.Printf("MCPING (%s):\n", addr)
	status, delay, err := mcping.PingAndList(addr, *protocol)
	if err != nil {
		exit(err)
	}

	fmt.Printf(`server: %s
protocol: %d
description: %s
delay: %v
list: %d/%d
`,
		status.Version.Name,
		status.Version.Protocol,
		status.Description, delay,
		status.Players.Online,
		status.Players.Max)
	for _, v := range status.Players.Sample {
		fmt.Printf("- [%s] %v\n", v.Name, v.ID)
	}
}

// 尝试SRV解析，若失败则使用默认端口
func lookupMC(addr string) (host, port string) {
	// TODO: 解析SRV
	return addr, "25565"
}

func exit(err error) {
	fmt.Printf("error: %v\n", err)
	os.Exit(1)
}
