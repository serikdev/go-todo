package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/serikdev/go-todo/internal/app/model"
	"github.com/serikdev/go-todo/internal/app/repository"
	"github.com/serikdev/go-todo/internal/app/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	repository.InitDB()
	routes.InitRoutes(r)
	return r
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()

	taskJSON := `{"title":"Test Task","description":"This is a test task","due_data":"2024-12-31"}`
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var task model.Task
	if err := json.Unmarshal(w.Body.Bytes(), &task); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if task.Title != "Test Task" {
		t.Errorf("Expected task title to be 'Test Task', but got '%s'", task.Title)
	}
}

func TestGetTask(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var tasks []model.Task
	if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(tasks) == 0 {
		t.Errorf("Expected to find tasks, but got none")
	}
}

func TestGetById(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var task model.Task
	if err := json.Unmarshal(w.Body.Bytes(), &task); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("Expected task ID to be 1, but got %d", task.ID)
	}
}
