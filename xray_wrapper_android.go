package libXray

import (
	"io"
	"os"
	"path/filepath"

	xrayfs "github.com/amnezia-vpn/amnezia-xray-core/common/platform/filesystem"
	mobasset "golang.org/x/mobile/asset"
)

// Prepare Xray for android
func InitXray() {
	xrayfs.NewFileReader = func(path string) (io.ReadCloser, error) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			_, file := filepath.Split(path)
			return mobasset.Open(file)
		}
		return os.Open(path)
	}
}
