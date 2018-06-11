package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"projekat/data"
	"projekat/dto"
	"projekat/serverErr"
)

var (
	CreateFilledTest = createFilledTest
	RemoveFilledTest = removeFilledTest
	GetFilledTest    = getFilledTest
	GetFilledTests   = getFilledTests
)

func createFilledTest(ctx context.Context, requestBody io.Reader) (err error) {
	request := &dto.CreateFilledRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	err = data.CreateFilled(ctx, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func removeFilledTest(ctx context.Context, filledTestUID string) (err error) {
	err = data.DeleteFilledTest(ctx, filledTestUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getFilledTest(ctx context.Context, filledTestUID string) (response *dto.GetFilledTestResponse, err error) {

	response, err = data.GetFilledTest(ctx, filledTestUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func getFilledTests(ctx context.Context) (response *dto.GetFilledTestsResponse, err error) {
	response, err = data.GetFilledTests(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}
