package ui

import (
	"embed"
)

// Special comment directive to embed the static files into the binary

//go:embed "html" "static"
var Files embed.FS
