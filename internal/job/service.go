package job

import (
	"time"

	"github.com/DanielTitkov/go-ent-echo-template/internal/app"
	"github.com/DanielTitkov/go-ent-echo-template/internal/configs"
	"github.com/DanielTitkov/go-ent-echo-template/internal/logger"
)

type Service struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewService(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Service {
	return &Service{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}

func (s *Service) SampleJob() {
	go func() {
		for {
			s.logger.Info("done sample job", "")
			time.Sleep(time.Duration(s.cfg.Job.SampleJobPeriod) * time.Second)
		}
	}()
}
