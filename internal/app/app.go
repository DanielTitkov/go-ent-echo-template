package app

import (
	"github.com/DanielTitkov/go-ent-echo-template/internal/configs"
	"github.com/DanielTitkov/go-ent-echo-template/internal/domain"
	"github.com/DanielTitkov/go-ent-echo-template/internal/logger"
)

type (
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		repo   Repository
	}
	Repository interface {
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserCount() (int, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}
}
