package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"projekat/data"
	"projekat/dto"
	"projekat/serverErr"
	"projekat/util"
)

var (
	CreateUser = createUser
	UpdateUser = updateUser
	RemoveUser = removeUser
	GetUser    = getUser
	GetUsers   = getUsers
)

func createUser(ctx context.Context, requestBody io.Reader) (response *dto.CreateUserResponse, err error) {
	request := &dto.CreateUserRequest{}
	response = &dto.CreateUserResponse{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}
	if request.EmployeeUID == "" {
		createPersonRequest := &dto.CreatePersonRequest{Address: request.Address, DateOfBirth: request.DateOfBirth, Email: request.Email, JMBG: request.JMBG, Name: request.Name, Surname: request.Surname}
		uid, err1 := data.CreatePerson(ctx, createPersonRequest)
		if err1 != nil {
			err = err1
			fmt.Println(err)
			return
		}
		personUID := uid
		createEmployeeRequest := &dto.CreateEmployeeRequest{PersonUID: personUID, RoleID: request.RoleID, WorkDocumentID: request.WorkDocumentID}
		uid, err1 = data.CreateEmployee(ctx, createEmployeeRequest)
		if err1 != nil {
			err = err1
			fmt.Println(err)
			return
		}
		request.EmployeeUID = uid
	}
	request.Password = util.GetMD5Hash(request.Password)
	uid, err := data.CreateUser(ctx, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	response.UID = uid
	return
}

func updateUser(ctx context.Context, userUID string, requestBody io.Reader) (err error) {
	request := &dto.UpdateUserRequest{}
	err = json.NewDecoder(requestBody).Decode(request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrBadRequest
		return
	}

	updateEmployee := &dto.UpdateEmployeeRequest{WorkDocumentID: request.WorkDocumentID}
	employeeUID, err := data.UpdateEmployeeForUser(ctx, userUID, updateEmployee)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}
	updatePerson := &dto.UpdatePersonRequest{Name: request.Name, Surname: request.Surname, JMBG: request.JMBG, Email: request.Email, Address: request.Address, DateOfBirth: request.DateOfBirth}
	err = data.UpdatePersonForEmployee(ctx, employeeUID, updatePerson)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}
	err = data.UpdateUser(ctx, userUID, request)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func removeUser(ctx context.Context, userUID string) (err error) {
	err = data.DeleteUser(ctx, userUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getUser(ctx context.Context, userUID string) (response *dto.GetUserResponse, err error) {

	response, err = data.GetUser(ctx, userUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func getUsers(ctx context.Context) (response *dto.GetUsersResponse, err error) {
	response, err = data.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}
	return
}
