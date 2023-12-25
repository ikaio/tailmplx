package main

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// IKAIO: env on top so USE_DATABASE is available for database.go
	_ "github.com/joho/godotenv/autoload"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/database"
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

		if os.Getenv("ENABLE_PUBLIC_FOLDER_CACHE") != "1" {
			// 1 week
			w.Header().Set("Cache-Control", "public, max-age=604800")
		}

		fs := http.StripPrefix(pathPrefix, http.FileServer(rootfs))
		fs.ServeHTTP(w, r)
	})
}

func main() {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	FileServer(r, "/", "./public")
	// NOTE: components and pages may overwrite files
	components.Init(r)
	database.Init()
	
	fmt.Println("[CONFIG] Choosen PORT:", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic(err)
	}
}
