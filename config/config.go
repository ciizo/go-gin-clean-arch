package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server Server
	Db     Database
}

type Server struct {
	Hostname string
	Port     int
}

type Database struct {
	ConnectionString string
}

const (
	configHostname   = "Hostname"
	configPort       = "Port"
	connectionString = "ConnectionString"
)

const (
	defaultPort = 8080
)

type env func(key string) string

type cfg struct {
	getEnv env
}

func New() *cfg {
	return &cfg{getEnv: os.Getenv}
}

func (c *cfg) LoadConfigurations() Config {
	return Config{
		Server: Server{
			Hostname: c.envString(configHostname, ""),
			Port:     c.envInt(configPort, defaultPort),
		},
		Db: Database{ConnectionString: c.envString(connectionString, "")},
	}
}

func (c *cfg) envString(key, defaultValue string) string {
	val := c.getEnv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (c *cfg) envInt(key string, defaultValue int) int {
	v := c.getEnv(key)

	val, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return val
}

func (c *cfg) envBool(key string, defaultValue bool) bool {
	v := c.getEnv(key)

	val, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}

	return val
}
