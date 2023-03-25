package gjwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UID      string    `json:"uid"`
	Role     string    `json:"role"`
	ExpireAt time.Time `json:"expire_at"`
	jwt.StandardClaims
}
