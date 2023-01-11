package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/h00s-go/tiny-link-backend/config"
)

type MemStore struct {
	Client *redis.Client
}

func NewMemStore(config *config.MemStore) *MemStore {
	return &MemStore{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
			Password: config.Password,
			DB:       config.Database,
		}),
	}
}

func (m *MemStore) Connect() error {
	if _, err := m.Client.Ping(context.Background()).Result(); err != nil {
		return err
	}
	return nil
}

func (m *MemStore) Close() {
	m.Client.Close()
}
