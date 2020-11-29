package handler

import (
	"net/http"

	"github.com/DanielTitkov/go-ent-echo-template/internal/api/model"
	"github.com/DanielTitkov/go-ent-echo-template/internal/api/util"
	"github.com/DanielTitkov/go-ent-echo-template/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) CreateUserHandler(c echo.Context) error {
	request := new(model.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	err := h.app.CreateUser(&domain.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "user created",
	})
}

func (h *Handler) GetUserHandler(c echo.Context) error {
	username, err := util.UsernameFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	u, err := h.app.GetUser(&domain.User{Username: username})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.GetUserResponse{
		Username: u.Username,
		Email:    u.Email,
	})
}

func (h *Handler) GetTokenHandler(c echo.Context) error {
	request := new(model.GetTokenRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	user := &domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	valid, err := h.app.ValidateUserPassword(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Message: "failed to authorize",
			Error:   err.Error(),
		})
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Message: "password is invalid",
		})
	}

	token, err := h.app.GetUserToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get token",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.GetTokenResponse{
		Token: token,
	})
}
