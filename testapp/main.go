// main.go

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gotuna/gotuna"
)

func main() {
	app := gotuna.App{
		ViewFiles: os.DirFS("."),
		Router:    gotuna.NewMuxRouter(),
	}
	
	app.Router.Handle("/", handlerHome(app))
	app.Router.Handle("/login", handlerLogin(app)).Methods(http.MethodGet, http.MethodPost)

	fmt.Println("Running on http://localhost:8888")
	http.ListenAndServe(":8888", app.Router)
}

func handlerHome(app gotuna.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NewTemplatingEngine().
			Render(w, r, "app.html")
	})
}

func handlerLogin(app gotuna.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login form...")
	})
}
