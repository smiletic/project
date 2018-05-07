package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"masterRad/data"
	"masterRad/dto"
	"masterRad/serverErr"
)

var (
	CreatePerson = createPerson
	UpdatePerson = updatePerson
	RemovePerson = removePerson
	GetPerson    = getPerson
	GetPersons   = getPersons
)

func createPerson(ctx context.Context, requestBody io.Reader) (response *dto.CreatePersonResponse, err error) {
	request := &dto.CreatePersonRequest{}
	response = &dto.CreatePersonResponse{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	uid, err := data.CreatePerson(ctx, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	response.UID = uid
	return
}

func updatePerson(ctx context.Context, personUID string, requestBody io.Reader) (err error) {
	request := &dto.UpdatePersonRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	err = data.UpdatePerson(ctx, personUID, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func removePerson(ctx context.Context, personUID string) (err error) {
	err = data.DeletePerson(ctx, personUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getPerson(ctx context.Context, personUID string) (response *dto.GetPersonResponse, err error) {

	response, err = data.GetPerson(ctx, personUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func getPersons(ctx context.Context) (response *dto.GetPersonsResponse, err error) {
	response, err = data.GetPersons(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}
