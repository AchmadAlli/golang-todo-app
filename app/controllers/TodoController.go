package controllers

import (
	"fmt"
	"net/http"

	"github.com/AchmadAlli/golang-todo-app/app/services"
	"github.com/AchmadAlli/golang-todo-app/app/utils"
	"github.com/gorilla/mux"
)

var service *services.TodoService

func ListenTodo(route *mux.Router, svc *services.TodoService) {
	fmt.Println("\nlistening todo controller")
	service = svc

	route.HandleFunc("/", index)
}

func index(writer http.ResponseWriter, req *http.Request) {
	todos, err := service.Index()

	if err != nil {
		utils.ErrorResponse(writer, 400, err.Error())
	}

	utils.Response(writer, http.StatusOK, todos)
}
