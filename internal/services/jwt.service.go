package services

import (
	"gotaskapp/internal/configs"
	"gotaskapp/internal/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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

func (s *jwtService) CreateToken(userId uint64) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.New().String()

	var err error
	atClamins := jwt.MapClaims{}
	atClamins["authorize"] = true
	atClamins["access_uuid"] = td.AccessUuid
	atClamins["user_id"] = userId
	atClamins["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClamins)
	td.AccessToken, err = at.SignedString([]byte(s.config.Jwt.Access))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte((s.config.Jwt.Refresh)))

	if err != nil {
		return nil, err
	}

	return td, nil
}
