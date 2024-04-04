package static

import "embed"

//go:embed all:dist
var Static embed.FS
