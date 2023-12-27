package router

import (
	"fmt"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()

	// A good base middleware stack
	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	Router.Use(middleware.Timeout(60 * time.Second))

	FileServer(Router, "/", "./public")

	fmt.Println("[CONFIG] PORT:", os.Getenv("PORT"))
	fmt.Println("[CONFIG] ENABLE_PUBLIC_FOLDER_CACHE:", os.Getenv("ENABLE_PUBLIC_FOLDER_CACHE"))
}
