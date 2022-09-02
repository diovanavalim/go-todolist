package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"todolist/cmd/api/database"
	"todolist/cmd/api/model"
	"todolist/cmd/api/response"
	"todolist/cmd/api/service"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var task model.Task

	if err = json.Unmarshal(requestBody, &task); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := task.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	taskService := service.CreateTaskService(db)

	result, err := taskService.CreateTask(task)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, struct {
		TaskID uint64 `json:"taskId"`
	}{
		TaskID: result,
	})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	taskService := service.CreateTaskService(db)

	task, err := taskService.GetTasks()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	taskService := service.CreateTaskService(db)

	result, err := taskService.GetTask(taskId)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var task model.Task

	if err = json.Unmarshal(requestBody, &task); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	taskService := service.CreateTaskService(db)

	if err := taskService.UpdateTask(taskId, task); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	taskService := service.CreateTaskService(db)

	if err = taskService.DeleteTask(taskId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
