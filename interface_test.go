package iputil

import (
	"net"
	"testing"
)

func TestInterfaceAddrs(t *testing.T) {
	ifAddrs, err := InterfaceAddrs()
	if err != nil {
		t.Error(err)
	}

	lo, err := net.InterfaceByIndex(1)
	if err != nil {
		t.Error(err)
	}
	loAddrs, err := lo.Addrs()
	if err != nil {
		t.Error(err)
	}

	if ifAddrs[0].String() != loAddrs[0].String() {
		t.Errorf("unexpected address: want %s, get %s", loAddrs[0], ifAddrs[0])
	}

	if ifAddrs[0].InterfaceName != lo.Name {
		t.Errorf("unexpected interface: want %s, get %s", ifAddrs[0].InterfaceName, lo.Name)
	}
}
