package models

/**
	工具
 */

import (
	"crypto/md5"
	"io"
	"fmt"
	"log"
)

func MD5(s string) string {
	h := md5.New()
	salt1 := "salt4shorturl"
	io.WriteString(h, s+salt1)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	log.Println(s, urlmd5)
	return urlmd5
}
