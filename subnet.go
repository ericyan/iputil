package iputil

import (
	"math/big"
	"net"
)

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

// Subnets divides the supernet into smaller subnets of given prefix
// size. It returns nil if subnet prefix size is invalid.
func Subnets(supernet *net.IPNet, prefix int) []*net.IPNet {
	ones, bits := supernet.Mask.Size()
	if ones > prefix || bits < prefix {
		return nil
	}

	mask := net.CIDRMask(prefix, bits)
	ip := new(big.Int).SetBytes(supernet.IP)
	size := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(bits-prefix)), big.NewInt(0))

	subnets := make([]*net.IPNet, 2<<uint(prefix-ones-1))
	for i := 0; i < len(subnets); i++ {
		subnets[i] = &net.IPNet{
			IP:   new(big.Int).Add(ip, new(big.Int).Mul(size, big.NewInt(int64(i)))).Bytes(),
			Mask: mask,
		}
	}

	return subnets
}
