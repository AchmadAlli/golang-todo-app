package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) InitRoutes() {
	if app.Router == nil {
		app.Router = mux.NewRouter()
	}
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
