// Package jwt provides JWT token creation and parsing.
//
// It wraps golang-jwt to provide typed claims with user identity fields
// (UserId, Username, NickName, UUID) and automatic expiration management.
package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mangk/adminBox/pkg/config"
	"github.com/mangk/adminBox/pkg/util"
)

// Claims represents the JWT claims with user identity information.
type Claims struct {
	UserId     int       `json:"user_id"`
	Username   string    `json:"username"`
	NickName   string    `json:"nick_name"`
	UUID       uuid.UUID `json:"uuid"`
	BufferTime int64     `json:"buffer_time"`
	jwt.StandardClaims
}

type Token struct {
	signingKey []byte
}

// New creates a new Token with the given signing key.
func New(signingKey []byte) *Token {
	return &Token{signingKey: signingKey}
}

// Create generates a signed JWT string from the provided claims and config.
func (j *Token) Create(user Claims, cfg config.JWT) (string, error) {
	bf, _ := util.ParseDuration(cfg.BufferTime)
	ep, _ := util.ParseDuration(cfg.ExpiresTime)
	user.BufferTime = int64(bf / time.Second)
	user.StandardClaims = jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", user.UserId),
		ExpiresAt: time.Now().Add(ep).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    cfg.Issuer,
		NotBefore: time.Now().Unix() - 1000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(j.signingKey)
}

// Parse verifies and extracts claims from a JWT string.
func (j *Token) Parse(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	t, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if t != nil && t.Valid {
		return claims, nil
	}
	return nil, errors.New("token parsing failed")
}
