package netutil

import (
	"errors"
	"fmt"
	"net"
)

func GetNextFree() (int, error) {
	for port := 49152; port <= 65535; port++ {
		_, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%v", port))
		if err == nil {
			return port, nil
		}
	}
	return 0, errors.New("No free port found")
}

func GetUsedPorts() []int {
	ports := []int{}
	for port := 49152; port <= 65535; port++ {
		adr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%v", port))
		if err != nil {
			continue
		}
		l, err := net.ListenTCP("tcp", adr)
		if err != nil {
			ports = append(ports, port)
		}
		defer l.Close()
	}
	return ports
}
