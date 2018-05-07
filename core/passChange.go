package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"masterRad/data"
	"masterRad/dto"
	"masterRad/serverErr"
	"masterRad/util"
)

var (
	ChangePass = changePass
)

func changePass(ctx context.Context, requestBody io.Reader, userUID string) (err error) {
	request := &dto.ChangePassRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	passed, err := data.CheckPassword(ctx, util.GetMD5Hash(request.OldPass), userUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}
	if !passed {
		err = serverErr.ErrNotAuthenticated
		return
	}
	err = data.ChangePassword(ctx, util.GetMD5Hash(request.NewPass), userUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}
	return
}
