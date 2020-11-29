package entgo

import (
	"context"

	"github.com/DanielTitkov/go-ent-echo-template/internal/domain"
	"github.com/DanielTitkov/go-ent-echo-template/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetUserByUsername(username string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (r *EntgoRepository) CreateUser(u *domain.User) (*domain.User, error) {
	user, err := r.client.User.
		Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPasswordHash(u.PasswordHash).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (r *EntgoRepository) GetUserCount() (int, error) {
	return r.client.User.Query().Count(context.Background())
}
