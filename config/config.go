package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ConfigEnv interface {

}

func New(env...string) ConfigEnv  {
	err := godotenv.Load(env...)
	if err != nil {
		panic(err)
	}
	return &config{}
}

type config struct {

}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}