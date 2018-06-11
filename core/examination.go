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
	CreateExamination = createExamination
	RemoveExamination = removeExamination
	GetExaminations   = getExaminations
)

func createExamination(ctx context.Context, requestBody io.Reader) (response *dto.CreateExaminationResponse, err error) {
	request := &dto.CreateExaminationRequest{}
	response = &dto.CreateExaminationResponse{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	uid, err := data.CreateExamination(ctx, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	response.UID = uid
	return
}

func removeExamination(ctx context.Context, examinationUID string) (err error) {
	err = data.DeleteExamination(ctx, examinationUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getExaminations(ctx context.Context) (response *dto.GetExaminationsResponse, err error) {
	response, err = data.GetExaminations(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}
