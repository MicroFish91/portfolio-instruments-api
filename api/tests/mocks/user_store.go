package mocks

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type MockUserStore struct{}

func NewMockUserStore() *MockUserStore {
	return &MockUserStore{}
}

func (s *MockUserStore) CreateUser(ctx context.Context, u types.User) (types.User, error) {
	return types.User{}, nil
}

func (s *MockUserStore) GetUserById(ctx context.Context, userId int) (types.User, error) {
	return types.User{}, nil
}

func (s *MockUserStore) GetUserByEmail(ctx context.Context, email string) (types.User, error) {
	return types.User{}, nil
}

func (s *MockUserStore) CreateSettings(ctx context.Context, settings types.Settings) (types.Settings, error) {
	return types.Settings{}, nil
}

func (s *MockUserStore) GetSettings(ctx context.Context, userId int) (types.Settings, error) {
	return types.Settings{}, nil
}

func (s *MockUserStore) UpdateSettings(ctx context.Context, settings types.Settings) (types.Settings, error) {
	return types.Settings{}, nil
}
