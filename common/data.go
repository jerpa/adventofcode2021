package common

import (
	"crypto/md5"
	"encoding/json"
)

// Checksum calculates a MD5 sum for given interface
func Checksum(data interface{}) [16]byte {
	jsonBytes, _ := json.Marshal(data)
	return md5.Sum(jsonBytes)
}
