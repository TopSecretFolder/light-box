package vue

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var embeddedDistFS embed.FS

// DistFS provides an http.FileSystem for the embedded 'dist' directory.
// It creates a sub-filesystem to serve files from the root of 'dist'
// rather than 'dist/index.html', etc.
func DistFS() http.FileSystem {
	// Create a sub-filesystem that makes the 'dist' directory the root.
	// The paths within embeddedDistFS will be "dist/index.html", "dist/css/style.css", etc.
	// fs.Sub makes it so that you access them as "index.html", "css/style.css".
	subFS, err := fs.Sub(embeddedDistFS, "dist")
	if err != nil {
		// This should ideally not happen if the embed directive and path are correct.
		log.Fatalf("editor: failed to create sub-filesystem for embedded dist: %v", err)
	}
	return http.FS(subFS)
}

// You can also provide the raw io/fs.FS if needed elsewhere, though for Echo's
// static serving, http.FileSystem is often more direct or what StaticWithConfig expects.
func RawDistFS() (fs.FS, error) {
	return fs.Sub(embeddedDistFS, "dist")
}
