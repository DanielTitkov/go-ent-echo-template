package prepare

import (
	"context"

	"github.com/DanielTitkov/go-ent-echo-template/internal/repository/entgo/ent"
	"github.com/DanielTitkov/go-ent-echo-template/internal/repository/entgo/ent/migrate"
)

func Migrate(ctx context.Context, client *ent.Client) error {
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}
	return nil
}
