package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	apiV1 := "/api/v1"

	mux.HandleFunc("/", s.HelloWorldHandler)

	AddListsHandlers(mux, s, apiV1)
	AddTasksHandlers(mux, s, apiV1)

	return mux
}

func AddListsHandlers(mux *http.ServeMux, s *Server, apiVersion string) {
	mux.HandleFunc("GET "+apiVersion+"/lists", s.GetListsHandler)
	mux.HandleFunc("POST "+apiVersion+"/lists", s.PostListsHandler)
	mux.HandleFunc("PUT "+apiVersion+"/lists/{id}", s.PutListHandler)
	mux.HandleFunc("DELETE "+apiVersion+"/lists/{id}", s.DeleteListHandler)
}

func AddTasksHandlers(mux *http.ServeMux, s *Server, apiVersion string) {
	mux.HandleFunc("GET "+apiVersion+"/lists/{listID}/tasks", s.GetTasksHandler)
	mux.HandleFunc("POST "+apiVersion+"/lists/{listID}/tasks", s.PostTasksHandler)
	mux.HandleFunc("PUT "+apiVersion+"/lists/{listID}/tasks/{taskID}", s.PutTaskHandler)
	mux.HandleFunc("DELETE "+apiVersion+"/lists/{listID}/tasks/{taskID}", s.DeleteTaskHandler)
	mux.HandleFunc("PATCH "+apiVersion+"/lists/{listID}/tasks/{taskID}/done", s.PatchTaskDoneHandler)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Hello World"

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
