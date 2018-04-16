package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"masterRad/data"
	"masterRad/dto"
	"masterRad/serverErr"
	"net/url"
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

func getPatients(ctx context.Context, queryParams url.Values) (response *dto.GetPatientsResponse, err error) {
	name := queryParams.Get("Name")
	surname := queryParams.Get("Surname")
	if name != "" || surname != "" {
		response, err = data.GetPatientsByName(ctx, name, surname)
		if err != nil {
			fmt.Println(err)
			err = serverErr.ErrInternal
		}
		return
	}
	medicalRecordUID := queryParams.Get("MedicalRecordId")
	healthCardUID := queryParams.Get("HealthCardId")
	if medicalRecordUID != "" || healthCardUID != "" {
		response, err = data.GetPatientsByHealthDocUIDs(ctx, medicalRecordUID, healthCardUID)
		if err != nil {
			fmt.Println(err)
			err = serverErr.ErrInternal
		}
		return
	}
	return
}
