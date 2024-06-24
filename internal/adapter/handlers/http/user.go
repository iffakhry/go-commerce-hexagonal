package http

import (
	"net/http"
	"strings"

	"github.com/iffakhry/go-commerce-hexagonal/internal/adapter/handlers/http/requests"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/pkg/webresponses"
	"github.com/iffakhry/go-commerce-hexagonal/internal/core/ports"
	"github.com/labstack/echo/v4"
)

// UserHandler represents the HTTP handler for user-related requests
type UserHandler struct {
	svc ports.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc ports.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// Register godoc
// @Summary      register users
// @Description  register users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request    body    requests.UserRequest  true  "data user"
// @Success      200  {object}   domain.User
// @Router       /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	// membaca data dari request body
	newUser := requests.UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, webresponses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}

	// mapping  dari request ke core
	inputCore := requests.UserRequestToEntity(newUser)
	// memanggil/mengirimkan data ke method service layer
	errInsert := h.svc.Create(inputCore)
	if errInsert != nil {

		// validasi code response
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, webresponses.JSONWebResponse("error insert data: "+errInsert.Error(), nil))

		}
		return c.JSON(http.StatusInternalServerError, webresponses.JSONWebResponse("error insert data: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, webresponses.JSONWebResponse("success add user", nil))
}
