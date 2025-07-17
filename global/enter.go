package global

import (
	"blogx_server/conf"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var Config *conf.Config
var DB *gorm.DB
var Redis *redis.Client
