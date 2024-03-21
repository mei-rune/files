package swaggerFiles

import (
	"github.com/swaggo/files/v2/resources"
)


// FS holds embedded swagger ui files
var FS, Err = resources.GetStaticDir()
