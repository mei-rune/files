package newversion

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var f embed.FS

func GetStaticDir() (fs.FS, error) {
	return fs.Sub(f, "dist")
}
