package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/util"
)

func NewJWT(signingKey []byte) *JWT {
	return &JWT{signingKey: signingKey}
}

type JWT struct {
	signingKey []byte
	Data       jwtUserInfo
}

type jwtUserInfo struct {
	UserId     int
	Username   string
	NickName   string
	UUID       uuid.UUID
	BufferTime int64
	jwt.StandardClaims
}

func (j *JWT) Create(user SysUser, config config.JWT) (string, error) {
	bf, _ := util.ParseDuration(config.BufferTime)
	ep, _ := util.ParseDuration(config.ExpiresTime)
	j.Data.UserId = user.ID
	j.Data.Username = user.Username
	j.Data.NickName = user.NickName
	j.Data.UUID = user.UUID
	j.Data.BufferTime = int64(bf / time.Second)
	j.Data.StandardClaims = jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", user.ID),
		ExpiresAt: time.Now().Add(ep).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    config.Issuer,
		NotBefore: time.Now().Unix() - 1000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Data)
	return token.SignedString(j.signingKey)
}

func (j *JWT) Parse(token string) (jud *jwtUserInfo, err error) {
	t, err := jwt.ParseWithClaims(token, &j.Data, func(t *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if t != nil {
		if j, ok := t.Claims.(*jwtUserInfo); ok && t.Valid {
			return j, nil
		}
		return nil, errors.New("token 解析失败")
	}

	return nil, errors.New("token 解析失败")
}
