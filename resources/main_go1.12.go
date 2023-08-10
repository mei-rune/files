//go:build !go1.16
// +build !go1.16
package resources

import (
	"io/fs"

	"github.com/swaggo/files/resources/oldversion"
)

func GetStaticDir() (fs.FS, error) {
	return oldversion.GetStaticDir()
}
