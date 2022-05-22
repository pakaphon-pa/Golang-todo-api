package services

import (
	"gotaskapp/internal/configs"
	"gotaskapp/internal/models"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type jwtService struct {
	redis  *redis.Client
	config configs.Config
}

func NewJwtService(redis *redis.Client, configs configs.Config) models.JwtServiceInterface {
	return &jwtService{
		redis:  redis,
		config: configs,
	}
}

func (s *jwtService) CreateAuth(userId uint64, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	errAccess := s.redis.Set(configs.Ctx, td.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err()

	if errAccess != nil {
		return errAccess
	}

	errRefresh := s.redis.Set(configs.Ctx, td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Err()

	if errRefresh != nil {
		return errRefresh
	}

	return nil
}
