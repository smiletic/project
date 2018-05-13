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
	CreatePatient = createPatient
	UpdatePatient = updatePatient
	RemovePatient = removePatient
	GetPatient    = getPatient
	GetPatients   = getPatients
)

func createPatient(ctx context.Context, requestBody io.Reader) (response *dto.CreatePatientResponse, err error) {
	request := &dto.CreatePatientRequest{}
	response = &dto.CreatePatientResponse{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}
	if request.PersonUID == "" {
		createPersonRequest := &dto.CreatePersonRequest{Address: request.Address, DateOfBirth: request.DateOfBirth, Email: request.Email, JMBG: request.JMBG, Name: request.Name, Surname: request.Surname}
		uid, err1 := data.CreatePerson(ctx, createPersonRequest)
		if err1 != nil {
			err = err1
			fmt.Println(err)
			return
		}
		request.PersonUID = uid
	}

	uid, err := data.CreatePatient(ctx, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	response.UID = uid
	return
}

func updatePatient(ctx context.Context, patientUID string, requestBody io.Reader) (err error) {
	request := &dto.UpdatePatientRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}
	updatePerson := &dto.UpdatePersonRequest{Name: request.Name, Surname: request.Surname, JMBG: request.JMBG, Email: request.Email, Address: request.Address, DateOfBirth: request.DateOfBirth}
	err = data.UpdatePersonForPatient(ctx, patientUID, updatePerson)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}
	err = data.UpdatePatient(ctx, patientUID, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func removePatient(ctx context.Context, patientUID string) (err error) {
	err = data.DeletePatient(ctx, patientUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getPatient(ctx context.Context, patientUID string) (response *dto.GetPatientResponse, err error) {

	response, err = data.GetPatient(ctx, patientUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func getPatients(ctx context.Context) (response *dto.GetPatientsResponse, err error) {
	response, err = data.GetPatients(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}
	return
}
