package config

import (
    "fmt"
    "log"

    "entgo.io/ent/dialect"
    "mystats-server/internal/ent"

    _ "github.com/lib/pq"
)

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func NewDatabaseConfig() *DatabaseConfig {
    return &DatabaseConfig{
        Host:     "localhost",
        Port:     "5433",
        User:     "postgres",
        Password: "postgres",
        DBName:   "mystats_db",
    }
}

func (c *DatabaseConfig) ConnectionString() string {
    return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.Host, c.Port, c.User, c.Password, c.DBName)
}

func NewEntClient(config *DatabaseConfig) *ent.Client {
    client, err := ent.Open(dialect.Postgres, config.ConnectionString())
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    return client
}
