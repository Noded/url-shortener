package main

import (
	"crypto/md5"
	"encoding/base64"
)

func ShortingURL(url string) string {
	hash := md5.Sum([]byte(url))                          // MD5 hash
	encoded := base64.URLEncoding.EncodeToString(hash[:]) // coding into base64
	return encoded[:6]                                    // returning first six symbols
}
