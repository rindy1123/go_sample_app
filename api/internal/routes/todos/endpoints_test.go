package todos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/api/internal/models"
	"github.com/api/test/factory"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func before() (*gin.Engine, *httptest.ResponseRecorder, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	db := models.SetupTestDB()
	TodoRouter{}.SetupEndpoints(r, db)
	w := httptest.NewRecorder()
	return r, w, db
}

func TestCreate(t *testing.T) {
	r, w, _ := before()

	reqBody := PostReqBody{
		Title:  "Test Todo",
		Status: models.TODO,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(reqBody)
	req, _ := http.NewRequest("POST", "/todos", &buf)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf(`Expected status code to be %v, got %v`, http.StatusOK, w.Code)
	}
	var res map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf(`Expected response to be a map, got %v`, w.Body.String())
	}
	id := res["id"]
	if id == "" {
		t.Fatalf(`Expected id to be non-empty, got %v`, id)
	}
}

func TestGet(t *testing.T) {
	r, w, db := before()

	todo := factory.Factory{}.CreateTodo(db)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/todos/%s", todo.Base.ID), nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf(`Expected status code to be %v, got %v`, http.StatusOK, w.Code)
	}
	var res GetTodo
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf(`Expected response to be a GetTodo, got %v`, w.Body.String())
	}
	if res.Title != todo.Title || res.Status != todo.Status {
		t.Fatalf(`Expected todo to be %+v, got %+v`, todo, res)
	}
}

func TestList(t *testing.T) {
	r, w, db := before()

	factory.Factory{}.CreateTodo(db)
	factory.Factory{}.CreateTodo(db)
	req, _ := http.NewRequest("GET", "/todos", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf(`Expected status code to be %+v, got %+v`, http.StatusOK, w.Code)
	}
	var res ListTodos
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf(`Expected response to be a ListTodos, got %+v`, w.Body.String())
	}
	if len(res.Todos) != 2 {
		t.Fatalf(`Expected 2 todos, got %+v`, len(res.Todos))
	}
}

func TestUpdate(t *testing.T) {
	r, w, db := before()

	todo := factory.Factory{}.CreateTodo(db)
	reqBody := UpdateReqBody{
		Title: "Update",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(reqBody)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/todos/%s", todo.Base.ID), &buf)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf(`Expected status code to be %v, got %v`, http.StatusOK, w.Code)
	}

	var res GetTodo
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf(`Expected response to be a GetTodo, got %+v`, w.Body.String())
	}
	if res.Title != "Update" || res.Status != todo.Status {
		t.Fatalf(`Expected todo to be %+v, got %+v`, todo, res)
	}
}

func TestDelete(t *testing.T) {
	r, w, db := before()

	todo := factory.Factory{}.CreateTodo(db)
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/todos/%s", todo.Base.ID), nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf(`Expected status code to be %v, got %v`, http.StatusOK, w.Code)
	}
}
