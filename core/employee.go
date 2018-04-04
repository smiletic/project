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
	CreateEmployee = createEmployee
	UpdateEmployee = updateEmployee
	RemoveEmployee = removeEmployee
	GetEmployee    = getEmployee
	GetEmployees   = getEmployees
)

func createEmployee(ctx context.Context, requestBody io.Reader) (response *dto.CreateEmployeeResponse, err error) {
	request := &dto.CreateEmployeeRequest{}
	response = &dto.CreateEmployeeResponse{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	uid, err := data.CreateEmployee(ctx, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	response.UID = uid
	return
}

func updateEmployee(ctx context.Context, employeeUID string, requestBody io.Reader) (err error) {
	request := &dto.UpdateEmployeeRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	err = data.UpdateEmployee(ctx, employeeUID, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func removeEmployee(ctx context.Context, employeeUID string) (err error) {
	err = data.DeleteEmployee(ctx, employeeUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getEmployee(ctx context.Context, employeeUID string) (response *dto.GetEmployeeResponse, err error) {

	response, err = data.GetEmployee(ctx, employeeUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func getEmployees(ctx context.Context) (response *dto.GetEmployeesResponse, err error) {
	response, err = data.GetEmployees(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}