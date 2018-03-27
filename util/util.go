package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"masterRad/data"
	"masterRad/dto"
)

var (
	Login      = login
	GetMD5Hash = getMD5Hash
)

func login(ctx context.Context, name, pass string) (autorizacija *dto.Autorizacija, err error) {
	pass = GetMD5Hash(pass)
	return data.Login(ctx, name, pass)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
