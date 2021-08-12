package controllers

import (
	"fmt"
	"net/http"

	"github.com/AchmadAlli/golang-todo-app/app/utils"
	"github.com/gorilla/mux"
)

func ListenTodo(route *mux.Router) {
	fmt.Println("\nlisten todo controller")
	route.HandleFunc("/", index)
}

func index(writer http.ResponseWriter, req *http.Request) {
	response := map[string]string{
		"data": "todo app succeed",
	}

	utils.Response(writer, response)
}
