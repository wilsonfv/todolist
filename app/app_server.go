package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wilsonfv/todolist/app/controller"
	"github.com/wilsonfv/todolist/app/dao"
	"github.com/wilsonfv/todolist/app/model"
	"log"
	"net/http"
)

var td = dao.TaskDao{}

func ListTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := controller.ListAll(&td)

	if err != nil {
		replyWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	replyWithJson(w, http.StatusOK, tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		replyWithError(w, http.StatusInternalServerError, "Invalid request")
		return
	}

	if err := controller.AddTask(&td, task); err != nil {
		replyWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	replyWithJson(w, http.StatusCreated, task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		replyWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := controller.DeleteTask(&td, task); err != nil {
		replyWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	replyWithJson(w, http.StatusOK, map[string]string{"result": "deleted"})
}

func replyWithError(w http.ResponseWriter, code int, msg string) {
	replyWithJson(w, code, map[string]string{"error": msg})
}

func replyWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	td.Server = "localhost:27017"
	td.Database = "task_db"
	td.Collection = "tasks"

	td.Connect()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/task", ListTask).Methods("GET")
	router.HandleFunc("/task", AddTask).Methods("POST")
	router.HandleFunc("/task", DeleteTask).Methods("DELETE")

	log.Println("Server starting at localhost:8181/task")
	if err := http.ListenAndServe(":8181", router); err != nil {
		log.Fatal(err)
	}
}
