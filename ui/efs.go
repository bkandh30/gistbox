package ui

import (
	"embed"
)

// Special comment directive to embed the static files into the binary

//go:embed "static"
var Files embed.FS
