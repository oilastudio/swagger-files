package swaggerFiles

import (
	"embed"
	"github.com/oilastudio/swagger-files/internal"
	"golang.org/x/net/webdav"
	"io/fs"
)

//go:embed dist
var dist embed.FS

func NewHandler() *webdav.Handler {
	static, _ := fs.Sub(dist, "dist")
	return &webdav.Handler{
		FileSystem: internal.NewWebDAVFileSystemFromFS(static),
		LockSystem: webdav.NewMemLS(),
	}
}
