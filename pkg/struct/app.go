package _struct

import (
	"go-server/pkg/adapter/redis"
	"go-server/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Db        *pgxpool.Pool
	Cache     redis.IRedis
	Config    *config.Schema
}
