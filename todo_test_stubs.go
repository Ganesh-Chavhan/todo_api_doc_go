package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TODO: Connect this helper to the actual Todo API handler.
func newTestHandler(t *testing.T) http.Handler {
	t.Helper()
	t.Skip("Connect this test to the Todo API handler")
	return http.NotFoundHandler()
}

type todoResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type todoListResponse struct {
	Todos []todoResponse `json:"todos"`
}

func request(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	res := httptest.NewRecorder()
	h.ServeHTTP(res, req)
	return res
}

func checkStatus(t *testing.T, res *httptest.ResponseRecorder, want int) {
	t.Helper()
	if res.Code != want {
		t.Fatalf("status = %d, want %d; body: %s", res.Code, want, res.Body.String())
	}
}

func decodeTodo(t *testing.T, res *httptest.ResponseRecorder) todoResponse {
	t.Helper()
	var todo todoResponse
	if err := json.Unmarshal(res.Body.Bytes(), &todo); err != nil {
		t.Fatalf("invalid todo response: %v", err)
	}
	return todo
}

func createTodo(t *testing.T, h http.Handler) todoResponse {
	t.Helper()
	res := request(h, http.MethodPost, "/todos", `{"name":"Test task"}`)
	checkStatus(t, res, http.StatusCreated)
	return decodeTodo(t, res)
}

// Positive: create a todo; status defaults to pending.
func TestCreateTodo(t *testing.T) {
	h := newTestHandler(t)
	res := request(h, http.MethodPost, "/todos", `{"name":"Test task"}`)
	checkStatus(t, res, http.StatusCreated)
	todo := decodeTodo(t, res)
	if todo.ID == 0 || todo.Name != "Test task" || todo.Status != "pending" {
		t.Fatalf("unexpected created todo: %+v", todo)
	}
}

// Positive: get an existing todo by ID.
func TestGetTodo(t *testing.T) {
	h := newTestHandler(t)
	created := createTodo(t, h)
	res := request(h, http.MethodGet, fmt.Sprintf("/todos/%d", created.ID), "")
	checkStatus(t, res, http.StatusOK)
}

// Positive: list todos.
func TestListTodos(t *testing.T) {
	h := newTestHandler(t)
	createTodo(t, h)
	res := request(h, http.MethodGet, "/todos", "")
	checkStatus(t, res, http.StatusOK)
	var list todoListResponse
	if err := json.Unmarshal(res.Body.Bytes(), &list); err != nil || len(list.Todos) == 0 {
		t.Fatalf("expected a non-empty todo list; body: %s", res.Body.String())
	}
}

// Positive: update a todo's name and status.
func TestUpdateTodo(t *testing.T) {
	h := newTestHandler(t)
	created := createTodo(t, h)
	path := fmt.Sprintf("/todos/%d", created.ID)
	res := request(h, http.MethodPut, path, `{"name":"Updated task","status":"completed"}`)
	checkStatus(t, res, http.StatusOK)
	todo := decodeTodo(t, res)
	if todo.Name != "Updated task" || todo.Status != "completed" {
		t.Fatalf("unexpected updated todo: %+v", todo)
	}
}

// Positive: delete an existing todo.
func TestDeleteTodo(t *testing.T) {
	h := newTestHandler(t)
	created := createTodo(t, h)
	res := request(h, http.MethodDelete, fmt.Sprintf("/todos/%d", created.ID), "")
	checkStatus(t, res, http.StatusNoContent)
}

// Negative: missing required name returns 400 Bad Request.
func TestCreateTodoMissingName(t *testing.T) {
	h := newTestHandler(t)
	res := request(h, http.MethodPost, "/todos", `{}`)
	checkStatus(t, res, http.StatusBadRequest)
}

// Negative: unsupported status returns 400 Bad Request.
func TestCreateTodoInvalidStatus(t *testing.T) {
	h := newTestHandler(t)
	res := request(h, http.MethodPost, "/todos", `{"name":"Test task","status":"started"}`)
	checkStatus(t, res, http.StatusBadRequest)
}

// Negative: requesting a todo that does not exist returns 404 Not Found.
func TestGetTodoNotFound(t *testing.T) {
	h := newTestHandler(t)
	res := request(h, http.MethodGet, "/todos/999999999", "")
	checkStatus(t, res, http.StatusNotFound)
}
