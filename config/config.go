package config

import (
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database Database
	MemStore MemStore
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type MemStore struct {
	Host     string
	Port     string
	Password string
	Database int
}

func NewConfig() *Config {
	c := NewConfigDefaults()
	if err := c.loadConfigFromFile("config.toml"); err != nil {
		log.Println("Unable to load config.toml, loaded defaults...")
	}
	c.applyEnvirontmentVariables()

	return c
}

func NewConfigDefaults() *Config {
	return &Config{}
}

func (c *Config) loadConfigFromFile(path string) error {
	if _, err := toml.DecodeFile(path, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) applyEnvirontmentVariables() {
	applyEnvirontmentVariable("DATABASE_HOST", &c.Database.Host)
	applyEnvirontmentVariable("DATABASE_PORT", &c.Database.Port)
	applyEnvirontmentVariable("DATABASE_USER", &c.Database.User)
	applyEnvirontmentVariable("DATABASE_PASSWORD", &c.Database.Password)
	applyEnvirontmentVariable("DATABASE_NAME", &c.Database.Name)

	applyEnvirontmentVariable("MEMSTORE_HOST", &c.MemStore.Host)
	applyEnvirontmentVariable("MEMSTORE_PORT", &c.MemStore.Port)
	applyEnvirontmentVariable("MEMSTORE_PASSWORD", &c.MemStore.Password)
	applyEnvirontmentVariable("MEMSTORE_DATABASE", &c.MemStore.Database)
}

func applyEnvirontmentVariable(key string, value interface{}) {
	if env, ok := os.LookupEnv(key); ok {
		switch v := value.(type) {
		case *string:
			*v = env
		case *bool:
			if env == "true" || env == "1" {
				*v = true
			} else if env == "false" || env == "0" {
				*v = false
			}
		case *int:
			if number, err := strconv.Atoi(env); err == nil {
				*v = number
			}
		}
	}
}
