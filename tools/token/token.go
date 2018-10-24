package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"strconv"
)

var SecretKey = []byte("qwertyuioasdfghjkzxcvbnmjhgfwerty")

/**
 * 生成token
 */
func Generator(uid int) string {

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["uid"] = strconv.Itoa(uid)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return ""
	}

	return tokenString
}

/**
 * 检测token
 */
func Verify(token string) (string, string) {
	return parse(token)
}

func parse(tk string)  (string, string) {
	if tk == "" {
		return "10000", ""
	}

	token, _ := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})

	//解析失败
	if token == nil {
		return "10000", ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	//检测过期
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return "10000", ""
	}

	if ok && token.Valid {
		return "0", claims["uid"].(string)
	} else {
		return "10000", ""
	}

	//无效
	return "10000", ""
}
