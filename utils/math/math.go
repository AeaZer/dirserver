package math

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func MD5Time(salt string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(time.Now().Format(time.RFC3339Nano) + salt))
	return hex.EncodeToString(hash.Sum(nil))
}
