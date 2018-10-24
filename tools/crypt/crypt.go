package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"math/rand"
	"time"
)

/*获取6位随机码*/
func RandCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int63n(100000000))

	randByte := []byte(code)
	randStr := ""

	for i := 0; i < 6; i++ {
		randStr += string(randByte[i])
	}

	return randStr
}

//MD5加密
func MD5(plainText string) string {
	h := md5.New()

	h.Write([]byte(plainText)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr) // 输出加密结果
}

func SHA1(plainText string) string {
	h := sha1.New()
	h.Write([]byte(plainText))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

//生成密码
func GeneratorPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		fmt.Println("Password.Generator", err)
	}

	return string(hashed[:])
}

//验证密码
func VerifyPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err == nil {
		return true
	}

	return false
}