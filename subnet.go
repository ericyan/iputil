package iputil

import "net"

// NetworkAddr returns the network address, which is also the beginning
// address, of the subnet.
func NetworkAddr(subnet *net.IPNet) net.IP {
	return subnet.IP
}

// BroadcastAddr returns the broadcast address, which is also the ending
// address, of the subnet.
func BroadcastAddr(subnet *net.IPNet) net.IP {
	n := len(subnet.IP)
	if n != len(subnet.Mask) {
		return nil
	}

	addr := make(net.IP, n)
	for i := 0; i < n; i++ {
		addr[i] = subnet.IP[i] | ^subnet.Mask[i]
	}

	return addr
}
