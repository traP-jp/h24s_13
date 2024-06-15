package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	repo *Repository
}

func NewHandlers(repo *Repository) *Handlers {
	return &Handlers{repo: repo}
}

type getMeResponse struct {
	ID string `json:"id"`
}

func (h *Handlers) GetMe(c echo.Context) error {
	id := c.Request().Header.Get("X-Forwarded-User")
	return c.JSON(http.StatusOK, &getMeResponse{
		ID: id,
	})
}

type getUserResponse struct {
	Groups []string `json:"groups"`
}

func (h *Handlers) GetUser(c echo.Context) error {
	// Get request params
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing user id")
	}

	// Get groups
	user, err := h.repo.GetUser(id)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	groups := user.Groups

	// Respond
	return c.JSON(http.StatusOK, &getUserResponse{
		Groups: groups,
	})
}

func (h *Handlers) GetUserRandomConnection(c echo.Context) error {
	// Get request params
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing user id")
	}
	countStr := c.QueryParam("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "count must be an integer")
	}
	if count <= 0 || count > 30 {
		return echo.NewHTTPError(http.StatusBadRequest, "count must be a positive integer less than 31")
	}

	// Get "friend of friend" connections
	connections := make([]string, 0, count)
	for range count {
		// TODO: implement me
		connections = append(connections, id)
	}

	// Respond
	return c.JSON(http.StatusOK, connections)
}

type userConnection struct {
	ID       string  `json:"id"`
	Strength float64 `json:"strength"`
}

func (h *Handlers) GetUserConnections(c echo.Context) error {
	// Get request params
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing user id")
	}

	// Get user connections
	connections := make([]*userConnection, 0)
	// TODO: implement me
	connections = append(connections, &userConnection{
		ID:       id,
		Strength: 100.0,
	})

	// Respond
	return c.JSON(http.StatusOK, connections)
}

const quizChoiceCount = 5

func (h *Handlers) GetQuiz(c echo.Context) error {
	// Get request params
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing user id")
	}

	// Choose random 5 people
	choices := make([]string, 0, quizChoiceCount)
	// TODO: implement me
	for range quizChoiceCount {
		choices = append(choices, id)
	}

	// Respond
	return c.JSON(http.StatusOK, choices)
}

func (h *Handlers) GetQuizAnswer(c echo.Context) error {
	// Get request params
	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing user id")
	}
	userAnswersStr := c.QueryParam("answers") // comma delimited
	userAnswers := strings.Split(userAnswersStr, ",")
	if len(userAnswers) != quizChoiceCount {
		return echo.NewHTTPError(http.StatusBadRequest, "answer must be of length 5, comma delimited")
	}

	// Get correct answers
	answers := make([]string, 0, quizChoiceCount)
	// TODO: implement me
	for range quizChoiceCount {
		answers = append(answers, id)
	}

	// Respond
	return c.JSON(http.StatusOK, answers)
}
