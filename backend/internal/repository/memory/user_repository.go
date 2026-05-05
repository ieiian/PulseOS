package memory

import (
	"context"
	"sync"

	"github.com/tse/PulseOS/backend/internal/domain/user"
)

type UserRepository struct {
	mu       sync.RWMutex
	profile  user.Profile
	settings user.Settings
	stats    user.Stats
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		profile: user.Profile{
			ID:            1,
			Name:          "测试用户",
			Age:           28,
			Gender:        "male",
			HeightCM:      175,
			WeightKG:      70,
			PrimaryGoal:   user.GoalFatLoss,
			SecondaryGoal: []user.Goal{user.GoalSleepRecover, user.GoalStressRelief},
			HealthFlags:   []string{},
		},
		settings: user.Settings{
			NotificationsEnabled:        true,
			StepPermissionGranted:       false,
			MicrophonePermissionGranted: false,
			SleepReminderEnabled:        true,
		},
		stats: user.Stats{
			CurrentStreak: 3,
			DaysTracked:   15,
		},
	}
}

func (r *UserRepository) SaveProfile(_ context.Context, profile user.Profile) user.Profile {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.profile = profile
	return r.profile
}

func (r *UserRepository) GetProfile(_ context.Context) user.Profile {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.profile
}

func (r *UserRepository) SaveSettings(_ context.Context, settings user.Settings) user.Settings {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.settings = settings
	return r.settings
}

func (r *UserRepository) GetSettings(_ context.Context) user.Settings {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.settings
}

func (r *UserRepository) GetStats(_ context.Context) user.Stats {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.stats
}
