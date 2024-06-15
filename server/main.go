package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getEnv(fallback string, lookupKeys ...string) string {
	for _, key := range lookupKeys {
		v, ok := os.LookupEnv(key)
		if ok {
			return v
		}
	}
	return fallback
}

func main() {
	// Prepare components
	repo, err := NewRepository()
	if err != nil {
		panic(err)
	}
	h := NewHandlers(repo)

	portStr := getEnv("8080", "PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("port must be an integer: " + err.Error())
	}

	// Prepare routes
	e := echo.New()
	api := e.Group("/api")
	{
		api.GET("/users/:id", h.GetUser)
		api.GET("/users/:id/random", h.GetUserRandomConnection)
		api.GET("/users/:id/connections", h.GetUserConnections)
		api.GET("/quiz/new", h.GetQuiz)
		api.GET("/quiz/answer", h.GetQuizAnswer)
	}

	// Serve
	err = e.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
