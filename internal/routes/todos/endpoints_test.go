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
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func before(t *testing.T) (*gin.Engine, *httptest.ResponseRecorder, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	db := models.SetupTestDB(t)
	TodoRouter{}.SetupEndpoints(r, db)
	w := httptest.NewRecorder()
	return r, w, db
}

func TestCreate(t *testing.T) {
	r, w, _ := before(t)

	reqBody := PostReqBody{
		Title:  "Test Todo",
		Status: models.TODO,
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(reqBody)
	req, _ := http.NewRequest("POST", "/todos", &buf)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var res map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.NotEmpty(t, res["id"])
}

func TestGet(t *testing.T) {
	r, w, db := before(t)

	todo := factory.Factory{}.CreateTodo(db)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/todos/%s", todo.Base.ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res GetTodo
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, todo.Title, res.Title)
	assert.Equal(t, todo.Status, res.Status)
}

func TestList(t *testing.T) {
	r, w, db := before(t)

	factory.Factory{}.CreateTodo(db)
	factory.Factory{}.CreateTodo(db)
	req, _ := http.NewRequest("GET", "/todos", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res ListTodos
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Len(t, res.Todos, 2)
}

func TestUpdate(t *testing.T) {
	r, w, db := before(t)

	todo := factory.Factory{}.CreateTodo(db)
	reqBody := UpdateReqBody{
		Title: "Update",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(reqBody)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/todos/%s", todo.Base.ID), &buf)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res GetTodo
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, "Update", res.Title)
	assert.Equal(t, todo.Status, res.Status)
}

func TestDelete(t *testing.T) {
	r, w, db := before(t)

	todo := factory.Factory{}.CreateTodo(db)
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/todos/%s", todo.Base.ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
