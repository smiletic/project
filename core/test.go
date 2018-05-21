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
	"net/http"
	"os"
	"strconv"
)

var (
	CreateTest = createTest
	RemoveTest = removeTest
	GetTests   = getTests
	GetTest    = getTest
)

func createTest(ctx context.Context, request *http.Request) (err error) {
	name := request.FormValue("name")
	specialty, err := strconv.Atoi(request.FormValue("specialty"))
	if err != nil {
		fmt.Println(err)
	}
	file, header, err := request.FormFile("fileUpload")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// copy example
	f, err := os.OpenFile("./"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)

	defer f.Close()
	defer os.Remove("./" + header.Filename)
	questions, err := util.ParseFile(ctx, "./"+header.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	marshaledQuestions, err := json.Marshal(questions)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = data.CreateTest(ctx, name, specialty, marshaledQuestions)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func removeTest(ctx context.Context, testUID string) (err error) {
	err = data.DeleteTest(ctx, testUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
	}

	return
}

func getTest(ctx context.Context, testUID string) (response *dto.GetTestResponse, err error) {

	response, err = data.GetTest(ctx, testUID)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}

func getTests(ctx context.Context) (response *dto.GetTestsResponse, err error) {
	response, err = data.GetTests(ctx)
	if err != nil {
		fmt.Println(err)
		err = serverErr.ErrInternal
		return
	}

	return
}
