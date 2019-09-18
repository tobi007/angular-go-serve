package bind

import (
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
	"strings"
)

type BinaryFileSystem struct {
	fs http.FileSystem
}

func (b *BinaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *BinaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func NewBinaryFileSystem(fs *assetfs.AssetFS) *BinaryFileSystem {
	return &BinaryFileSystem {
		fs,
	}
}

