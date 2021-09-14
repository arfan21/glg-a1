package config

import (
	"os"

	"github.com/arfan21/hacktiv8-assignment-2/helper"
	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type config struct{}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	helper.PanicIfNeeded(err)
	return &config{}
}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}
