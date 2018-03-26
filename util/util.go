package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"masterRad/data"
)

var (
	Login      = login
	GetMD5Hash = getMD5Hash
)

func login(ctx context.Context, name, pass string) (existingUser bool) {
	pass = GetMD5Hash(pass)
	return data.CheckLogin(ctx, name, pass)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
