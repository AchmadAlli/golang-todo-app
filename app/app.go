package app

import (
	"log"
	"net/http"

	utils "github.com/AchmadAlli/golang-todo-app/app/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	db     *gorm.DB
}

func (app *App) InitRoutes() {
	if app.Router == nil {
		app.Router = mux.NewRouter()
	}
	app.Router.HandleFunc("/check", healthCheck)
}

func (app *App) AddDatabase(db *gorm.DB) {
	app.db = db
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}

func healthCheck(writer http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"data": "Ok",
	}

	utils.Response(writer, response)
}
