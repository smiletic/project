package util

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"projekat/config"
	"projekat/data"
	"projekat/dto"
	"projekat/enum"
	"projekat/serverErr"
	"strings"

	"github.com/tealeg/xlsx"
)

var (
	Login         = login
	GetMD5Hash    = getMD5Hash
	ParseFile     = parseFile
	GenerateID    = generateID
	CreateSession = createSession
)

var (
	UserKey = "UserKey"
)

func login(ctx context.Context, name, pass string) (autorizacija *dto.Authorization, err error) {
	pass = GetMD5Hash(pass)
	return data.Login(ctx, name, pass)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func createSession(ctx context.Context, userUID string) (authToken string, err error) {
	token := GenerateID()
	return token, data.CreateSession(ctx, userUID, token)
}

func parseFile(ctx context.Context, filePath string) (parsedQuestions *dto.TestQuestions, err error) {
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in ", r)
			err = serverErr.ErrBadRequest
			return
		}
	}()
	questions := &dto.TestQuestions{}
	questions.Questions = make([]*dto.Question, 0, 0)
	sheet := file.Sheets[0]
	var question *dto.Question
	for _, row := range sheet.Rows {
		if len(row.Cells) == 0 {
			break
		}
		switch strings.ToLower(row.Cells[0].String()) {
		case strings.ToLower(config.GetQuestionStartString()):
			question = &dto.Question{}
		case strings.ToLower(config.GetQuestionEndString()):
			questions.Questions = append(questions.Questions, question)
		case strings.ToLower(config.GetQuestionTextString()):
			question.Question = row.Cells[1].String()
		case strings.ToLower(config.GetQuestionTypeString()):
			question.Type = typeCodeFromName(row.Cells[1].String())
		case strings.ToLower(config.GetQuestionAnswersString()):
			question.Answers = make([]string, 0, 0)
			for i := 1; i < len(row.Cells); i++ {
				question.Answers = append(question.Answers, row.Cells[i].String())
			}
		}
	}
	parsedQuestions = questions
	return
}

func typeCodeFromName(name string) (code enum.QuestionType) {
	switch strings.ToLower(name) {
	case strings.ToLower(config.GetQuestionTypeNamesFreeText()):
		code = enum.QuestionTypeFreeText
	case strings.ToLower(config.GetQuestionTypeNamesFreeNumerical()):
		code = enum.QuestionTypeFreeNumerical
	case strings.ToLower(config.GetQuestionTypeNamesRadioGroup()):
		code = enum.QuestionTypeRadioGroup
	case strings.ToLower(config.GetQuestionTypeNamesCheckbox()):
		code = enum.QuestionTypeCheckbox
	default:
		code = enum.QuestionTypeFreeText
	}
	return
}

func generateID() string {
	raw := make([]byte, 18, 18)
	enc := make([]byte, 24, 24)

	rand.Read(raw)
	base64.RawURLEncoding.Encode(enc, raw)

	return string(enc)
}
