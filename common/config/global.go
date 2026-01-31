package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	pb "medi-biz/proto/medi"
)

var DB *gorm.DB
var ConfigData Config
var Ctx = context.Background()
var Rdb *redis.Client
var MediClient pb.MediClient
