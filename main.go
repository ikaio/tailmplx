package main

import (

	"net/http"
	"os"
	
	// Load ENV variables and sets up the DB var and Snowflake id generation
	_ "github.com/ikaio/tailmplx/database"
	// Sets up the Router var
	"github.com/ikaio/tailmplx/router"
	// IMPORTANT: Components needs to load after the router as they make use of init()
	_ "github.com/ikaio/tailmplx/components"
)

func main() {
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router.Router)
	if err != nil {
		panic(err)
	}
}
