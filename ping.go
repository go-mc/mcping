package mcping

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// PingAndList check server status and list online player.
// Returns server status and delay.
// The addr has the form "host:port". The port are required.
func PingAndList(addr string, protocol int) (*Status, time.Duration, error) {
	conn, err := mcnet.DialMC(addr)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: dial fail: %v", err)
	}
	return pingAndList(addr, conn, protocol)
}

// PingAndListConn is the version of PingAndList using a exist connection.
func PingAndListConn(conn net.Conn, protocol int) (*Status, time.Duration, error) {
	addr := conn.RemoteAddr().String()
	mcConn := mcnet.WrapConn(conn)
	return pingAndList(addr, mcConn, protocol)
}

func pingAndList(addr string, conn *mcnet.Conn, protocol int) (*Status, time.Duration, error) {
	// parse hostname and port
	host, strPort, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, 0, fmt.Errorf("could not split host and port: %v", err)
	}

	port, err := strconv.ParseUint(strPort, 10, 16)
	if err != nil {
		return nil, 0, fmt.Errorf("port must be a number: %v", err)
	}

	// handshake
	err = conn.WritePacket(pk.Marshal(0x00, // packet ID
		pk.VarInt(protocol),    // protocol version
		pk.String(host),        // server host
		pk.UnsignedShort(port), // server port
		pk.Byte(1),             // next: ping
	))
	if err != nil {
		return nil, 0, fmt.Errorf("sending handshake: %v", err)
	}

	// list
	err = conn.WritePacket(pk.Marshal(0))
	if err != nil {
		return nil, 0, fmt.Errorf("sending list: %v", err)
	}

	// response
	var recv pk.Packet
	err = conn.ReadPacket(&recv)
	if err != nil {
		return nil, 0, fmt.Errorf("receiving response: %v", err)
	}

	var s pk.String
	if err = recv.Scan(&s); err != nil {
		return nil, 0, fmt.Errorf("scanning list: %v", err)
	}

	// ping
	startTime := time.Now()
	unixStartTime := pk.Long(startTime.Unix())

	err = conn.WritePacket(pk.Marshal(0x01, unixStartTime))
	if err != nil {
		return nil, 0, fmt.Errorf("sending ping: %v", err)
	}

	err = conn.ReadPacket(&recv)
	if err != nil {
		return nil, 0, fmt.Errorf("receiving pong: %v", err)
	}
	delay := time.Since(startTime)

	var t pk.Long
	if err = recv.Scan(&t); err != nil {
		return nil, 0, fmt.Errorf("scanning pong: %v", err)
	}
	// check time
	if t != unixStartTime {
		return nil, 0, errors.New("mismatched pong")
	}

	// parse status
	status := new(Status)
	if err = json.Unmarshal([]byte(s), status); err != nil {
		return nil, 0, fmt.Errorf("unmarshal json fail: %v", err)
	}

	return status, delay, nil
}
