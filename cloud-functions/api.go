package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

// Todo 待办事项模型
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	todoMu    sync.RWMutex
	todoSeq   = 3
	todoStore = []Todo{
		{ID: 1, Title: "Deploy to EdgeOne", Completed: true, CreatedAt: time.Now().Add(-48 * time.Hour)},
		{ID: 2, Title: "Write Go handlers", Completed: true, CreatedAt: time.Now().Add(-24 * time.Hour)},
		{ID: 3, Title: "Add Echo framework", Completed: false, CreatedAt: time.Now()},
	}
)

func main() {
	e := echo.New()
	e.HideBanner = true

	// 中间件
	e.Use(panicRecover)
	e.Use(requestLogger)

	// 路由
	e.GET("/", welcome)
	e.GET("/health", health)

	// Todo CRUD
	api := e.Group("/api")
	api.GET("/todos", listTodos)
	api.POST("/todos", createTodo)
	api.GET("/todos/:id", getTodo)
	api.PATCH("/todos/:id/toggle", toggleTodo)
	api.DELETE("/todos/:id", deleteTodo)

	port := "9000"
	fmt.Printf("Echo server starting on :%s\n", port)
	s := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Fatal(e.StartServer(s))
}

// welcome GET /
func welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome to EdgeOne Echo Demo!",
		"version": "1.0.0",
		"routes": []string{
			"GET  /           - this page",
			"GET  /health     - health check",
			"GET  /api/todos  - list todos",
			"POST /api/todos  - create todo",
			"GET  /api/todos/:id        - get todo",
			"PATCH /api/todos/:id/toggle - toggle todo",
			"DELETE /api/todos/:id       - delete todo",
		},
	})
}

// health GET /health
func health(c echo.Context) error {
	hostname, _ := os.Hostname()
	return c.JSON(http.StatusOK, echo.Map{
		"status":    "ok",
		"framework": "echo",
		"goVersion": runtime.Version(),
		"os":        runtime.GOOS,
		"arch":      runtime.GOARCH,
		"hostname":  hostname,
	})
}

// listTodos GET /api/todos
func listTodos(c echo.Context) error {
	todoMu.RLock()
	defer todoMu.RUnlock()
	return c.JSON(http.StatusOK, echo.Map{"data": todoStore, "total": len(todoStore)})
}

// createTodo POST /api/todos
func createTodo(c echo.Context) error {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}
	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "title is required"})
	}
	todoMu.Lock()
	todoSeq++
	todo := Todo{ID: todoSeq, Title: req.Title, Completed: false, CreatedAt: time.Now()}
	todoStore = append(todoStore, todo)
	todoMu.Unlock()
	return c.JSON(http.StatusCreated, echo.Map{"data": todo})
}

// getTodo GET /api/todos/:id
func getTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	todoMu.RLock()
	defer todoMu.RUnlock()
	for _, t := range todoStore {
		if t.ID == id {
			return c.JSON(http.StatusOK, echo.Map{"data": t})
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
}

// toggleTodo PATCH /api/todos/:id/toggle
func toggleTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	todoMu.Lock()
	defer todoMu.Unlock()
	for i := range todoStore {
		if todoStore[i].ID == id {
			todoStore[i].Completed = !todoStore[i].Completed
			return c.JSON(http.StatusOK, echo.Map{"data": todoStore[i]})
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
}

// deleteTodo DELETE /api/todos/:id
func deleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	todoMu.Lock()
	defer todoMu.Unlock()
	for i, t := range todoStore {
		if t.ID == id {
			todoStore = append(todoStore[:i], todoStore[i+1:]...)
			return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "todo not found"})
}

// requestLogger 请求日志中间件
func requestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		fmt.Printf("[%s] %s %s → %d (%s)\n",
			time.Now().Format("15:04:05"),
			c.Request().Method,
			c.Request().URL.Path,
			c.Response().Status,
			time.Since(start).Round(time.Microsecond),
		)
		return err
	}
}

// panicRecover panic 恢复中间件
func panicRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("PANIC recovered: %v\n%s\n", r, debug.Stack())
				c.JSON(http.StatusInternalServerError, echo.Map{
					"error":   "Internal Server Error",
					"message": fmt.Sprintf("%v", r),
				})
			}
		}()
		return next(c)
	}
}
