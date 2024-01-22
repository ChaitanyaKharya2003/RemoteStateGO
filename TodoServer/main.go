package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

type todo struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
}

var todos = make([]todo, 0)

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "health\n")
}

func todoList(w http.ResponseWriter, req *http.Request) {
	RespondJSON(w, http.StatusOK, todos)
}

func addTodo(w http.ResponseWriter, req *http.Request) {
	var body todo
	if err := ParseBody(req.Body, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body.CreatedAt = time.Now()
	body.Id = uuid.New()
	todos = append(todos, body)
	w.WriteHeader(http.StatusCreated)
}

func updateTODO(w http.ResponseWriter, req *http.Request) {
	var body todo
	if err := ParseBody(req.Body, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i := range todos {
		if todos[i].Id == body.Id {
			todos[i] = body
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func completeTODO(w http.ResponseWriter, req *http.Request) {
	var body todo
	if err := ParseBody(req.Body, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i := range todos {
		if todos[i].Id == body.Id {
			todos[i].CompletedAt = time.Now()
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteTODO(w http.ResponseWriter, req *http.Request) {
	var body todo
	if err := ParseBody(req.Body, &body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newTodos := make([]todo, 0)
	for i := range todos {
		if todos[i].Id != body.Id {
			newTodos = append(newTodos, todos[i])
		}
	}
	todos = newTodos
	w.WriteHeader(http.StatusCreated)
}

func main() {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Get("/health", health)
		todoRouters(r)
	})
	fmt.Println("Server starting")
	http.ListenAndServe(":8090", router)
}
func todoRouters(r chi.Router) chi.Router {
	return r.Route("/todo", func(todoRouter chi.Router) {
		todoRouter.Get("/", todoList)
		todoRouter.Post("/add-todo", addTodo)
		todoIDRouters(todoRouter)
	})
}

func todoIDRouters(todoRouter chi.Router) chi.Router {
	return todoRouter.Route("/{id}", func(todoIDRouter chi.Router) {
		todoRouter.Put("/", updateTODO)
		todoRouter.Put("/complete", completeTODO)
		todoRouter.Delete("/", deleteTODO)
	})
}

// RespondJSON sends the rateMetricInterface as a JSON
func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		if err := EncodeJSONBody(w, body); err != nil {
			fmt.Println(fmt.Errorf("failed to respond JSON with error: %+v", err))
		}
	}
}

// EncodeJSONBody writes the JSON body to response writer
func EncodeJSONBody(resp http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(resp).Encode(data)
}

// ParseBody parses the values from io reader to a given interface
func ParseBody(body io.Reader, out interface{}) error {
	err := json.NewDecoder(body).Decode(out)
	if err != nil {
		return err
	}

	return nil
}
