//go:build unix

package libXray

import (
	"errors"
	"github.com/amnezia-vpn/amnezia-libxray/nodep"
	"github.com/amnezia-vpn/amnezia-tun2socks/v2/engine"
	"golang.org/x/sys/unix"
)

type Tun2SocksConfig engine.Key

// Start tun2socks engine
func StartTun2Socks(conf *Tun2SocksConfig, fd int) string {
	engine.Insert((*engine.Key)(conf))
	err := engine.StartTun2Socks()
	if err != nil {
		e := unix.Close(fd)
		err = errors.Join(err, e)
	}
	return nodep.WrapError(err)
}

// Stop tun2socks engine
func StopTun2Socks() string {
	err := engine.StopTun2Socks()
	return nodep.WrapError(err)
}
