package service

import (
	"context"

	"github.com/tse/PulseOS/backend/internal/domain/user"
)

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Onboard(ctx context.Context, profile user.Profile) user.Profile {
	return s.repo.SaveProfile(ctx, profile)
}

func (s *UserService) GetProfile(ctx context.Context) user.Profile {
	return s.repo.GetProfile(ctx)
}

func (s *UserService) UpdateProfile(ctx context.Context, profile user.Profile) user.Profile {
	return s.repo.SaveProfile(ctx, profile)
}

func (s *UserService) GetSettings(ctx context.Context) user.Settings {
	return s.repo.GetSettings(ctx)
}

func (s *UserService) UpdateSettings(ctx context.Context, settings user.Settings) user.Settings {
	return s.repo.SaveSettings(ctx, settings)
}

func (s *UserService) GetStats(ctx context.Context) user.Stats {
	return s.repo.GetStats(ctx)
}

