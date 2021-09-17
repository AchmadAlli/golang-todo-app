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
	DB     *gorm.DB
}

func (app *App) InitRoutes() {
	if app.Router == nil {
		app.Router = mux.NewRouter().PathPrefix("/api").Subrouter()
	}
	app.Router.HandleFunc("/check", healthCheck).Methods("GET")
}

func (app *App) AddDatabase(db *gorm.DB) {
	app.DB = db
}

func (app *App) Run() {
	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}

func healthCheck(writer http.ResponseWriter, r *http.Request) {
	utils.Response(writer, http.StatusOK, "OK")
}
