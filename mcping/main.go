package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	"github.com/go-mc/mcping"
	"github.com/mattn/go-colorable"
)

var protocol = flag.Int("p", 578, "The minecraft protocol version")
var favicon = flag.String("f", "", "Path to output server's icon")
var output = colorable.NewColorableStdout()

func main() {
	flag.Parse()
	addrs := lookupMC(flag.Arg(0))

	for _, addr := range addrs {
		fmt.Fprintf(output, "MCPING (%s):\n", addr)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Fprintf(output, "dial error: %v\n", err)
			continue
		}
		status, delay, err := mcping.PingAndListConn(conn, *protocol)
		if err != nil {
			fmt.Fprintf(output, "error: %v\n", err)
			continue
		}

		fmt.Fprintf(output,
			`	server: %s
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

		if *favicon != "" {
			fmt.Fprintf(output, "Save server's icon into %s\n", *favicon)
			icon, err := status.Favicon.ToPNG()
			if err != nil {
				fmt.Fprintf(output, "error: %v\n", err)
				continue
			}
			err = ioutil.WriteFile(*favicon, icon, 0666)
			if err != nil {
				fmt.Fprintf(output, "error: %v\n", err)
				continue
			}
		}
	}
}

// written after read mojang's code
func lookupMC(addr string) (addrs []string) {
	if !strings.Contains(addr, ":") {
		_, addrsSRV, err := net.LookupSRV("minecraft", "tcp", addr)
		if err == nil && len(addrsSRV) > 0 {
			for _, addrSRV := range addrsSRV {
				addrs = append(addrs, net.JoinHostPort(addrSRV.Target, strconv.Itoa(int(addrSRV.Port))))
			}
			return
		}
		return []string{net.JoinHostPort(addr, "25565")}
	}
	return []string{addr}
}
