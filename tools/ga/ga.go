package ga

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
	"strconv"
	"autumn/tools/crypt"
	"log"
)

// ga 获取key对应的验证码
func Code(key string) string {
	return generator(key, time.Now().Unix())
}

// ga 获取秘钥
func Secret() string {
	tstr := strconv.FormatInt(time.Now().UnixNano(), 10)
	return base32.StdEncoding.EncodeToString([]byte(crypt.RandCode() + tstr))
}

// ga 获取key and t对应的验证码
// key 秘钥
// t 1970年的秒
func generator(key string, t int64) string {
	hs, err := hmac_sha1(key, t/30)
	if err != nil {
		log.Println("tools.ga.generator:", err)
		return ""
	}

	snum := last_4byte(hs)
	d := snum % 1000000
	return fmt.Sprintf("%06d", d)
}


func last_4byte(hmacSha1 []byte) int32 {
	if len(hmacSha1) != sha1.Size {
		return 0
	}

	offsetBits := int8(hmacSha1[len(hmacSha1)-1]) & 0x0f
	p := (int32(hmacSha1[offsetBits]) << 24) | (int32(hmacSha1[offsetBits+1]) << 16) | (int32(hmacSha1[offsetBits+2]) << 8) | (int32(hmacSha1[offsetBits+3]) << 0)
	return (p & 0x7fffffff)
}

func hmac_sha1(key string, t int64) ([]byte, error) {
	decodeKey, err := base32.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	cData := make([]byte, 8)
	binary.BigEndian.PutUint64(cData, uint64(t))

	h1 := hmac.New(sha1.New, decodeKey)
	_, e := h1.Write(cData)
	if e != nil {
		return nil, e
	}
	return h1.Sum(nil), nil
}
