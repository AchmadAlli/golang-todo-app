package app

import (
	"log"
	"net/http"

	utils "github.com/AchmadAlli/golang-todo-app/app/utils"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) InitRoutes() {
	if app.Router == nil {
		app.Router = mux.NewRouter()
	}
	app.Router.HandleFunc("/check", checkHealth)
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}

func checkHealth(writer http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"data": "Ok",
	}

	utils.Response(writer, response)
}
