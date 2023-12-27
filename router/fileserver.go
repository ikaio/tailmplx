package router

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

// https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
// adapted for not found handling
func FileServer(router *chi.Mux, path, root string) {
	rootfs := http.Dir(root)

	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")

		_, err := os.Stat(filepath.Join(root, r.URL.Path))
		if errors.Is(err, fs.ErrNotExist) {
			router.NotFoundHandler().ServeHTTP(w, r)
			return
		}

		if os.Getenv("ENABLE_PUBLIC_FOLDER_CACHE") == "1" {
			// 1 week
			w.Header().Set("Cache-Control", "public, max-age=604800")
		}

		fs := http.StripPrefix(pathPrefix, http.FileServer(rootfs))
		fs.ServeHTTP(w, r)
	})
}
