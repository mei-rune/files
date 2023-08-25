//go:build !go1.16
// +build !go1.16
package resources

import (
	"io/fs"

	"github.com/swaggo/files/v2/resources/oldversion"
)

func GetStaticDir() (fs.FS, error) {
	return oldversion.GetStaticDir()
}
