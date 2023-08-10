//go:build go1.16
// +build go1.16

package resources

import (
	"io/fs"

	"github.com/swaggo/files/resources/newversion"
)

func GetStaticDir() (fs.FS, error) {
	return newversion.GetStaticDir()
}
