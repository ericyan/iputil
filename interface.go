package iputil

import (
	"net"
)

// InterfaceAddr represents an interface IP address and the name of its
// associated network interface.
type InterfaceAddr struct {
	*net.IPNet
	InterfaceName string
}

// Interface returns the associated network interface.
func (addr *InterfaceAddr) Interface() (*net.Interface, error) {
	return net.InterfaceByName(addr.InterfaceName)
}

// InterfaceAddrs returns a list of the system's unicast interface IP
// addresses.
func InterfaceAddrs() ([]*InterfaceAddr, error) {
	var ifAddrs []*InterfaceAddr

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			ifAddrs = append(ifAddrs, &InterfaceAddr{addr.(*net.IPNet), iface.Name})
		}
	}

	return ifAddrs, nil
}
