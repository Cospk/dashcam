package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	Config   config.Server
	Log      *zap.Logger
	SugarLog *zap.SugaredLogger
	DB       *gorm.DB
	Redis    *redis.Client
)
