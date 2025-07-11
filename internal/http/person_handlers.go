package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"person-crud/internal/models"
)

type PersonLogic interface {
	Create(p *models.Person) error
	GetAll() ([]*models.Person, error)
	GetByID(id int) (*models.Person, error)
	Update(id int, p *models.Person) error
	Delete(id int) error
}

type PersonHandler struct {
	logic PersonLogic
}

func NewPersonHandler(l PersonLogic) *PersonHandler {
	return &PersonHandler{logic: l}
}

func (h *PersonHandler) CreatePerson(c echo.Context) error {
	p := new(models.Person)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.logic.Create(p); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *PersonHandler) GetPersons(c echo.Context) error {
	people, err := h.logic.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, people)
}

func (h *PersonHandler) GetPerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	p, err := h.logic.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if p == nil {
		return echo.NewHTTPError(http.StatusNotFound, "person not found")
	}
	return c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) UpdatePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	p := new(models.Person)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.logic.Update(id, p); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) DeletePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	if err := h.logic.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
