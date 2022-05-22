package models

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type JwtServiceInterface interface {
	CreateAuth(userId uint64, td *TokenDetails) error
	CreateToken(userId uint64) (*TokenDetails, error)
}
