package secret

import (
	"fmt"
	"time"

	"github.com/YogeLiu/CloudDisk/pkg/conf"
	"github.com/golang-jwt/jwt/v4"
)

func EnJWT(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	ret, err := token.SignedString([]byte(conf.SystemConfig.Secret))
	if err != nil {
		return "-1"
	}
	return ret
}

func DeJwt(str string) int {
	token, err := jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return conf.SystemConfig.Secret, nil
	})
	if err != nil {
		return 0
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims.VerifyExpiresAt(time.Now().Unix(), false) {
		return claims["id"].(int)
	}
	return 0
}
