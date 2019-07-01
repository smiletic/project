package core

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"projekat/config"
	"projekat/data"
	"projekat/dto"
	"projekat/enum"
	"projekat/logger"
	"projekat/serverErr"
	"projekat/utils"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var (
	Login         = login
	CreateSession = createSession
	GetSession    = getSession
	RemoveSession = removeSession
	generateToken = _generateToken
)

const (
	UserKey = "UserKey"

	allowedTimeDifference = 5 * time.Minute
)

func login(ctx context.Context, name, pass string) (*dto.SessionInfo, error) {
	pass = utils.GetMD5Hash(pass)
	return data.Login(ctx, name, pass)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func createSession(ctx context.Context, userUID string) (token string, err error) {
	token = generateToken()
	err = data.CreateSession(ctx, userUID, token)
	if err != nil {
		logger.Error("Couldn't create session: %v", err)
		return
	}
	return
}

func getSession(ctx context.Context, authorizationString, userUID string) (*dto.SessionInfo, error) {
	sessions, err := data.GetSessionsForUser(ctx, userUID)
	if err != nil {
		logger.Error("Couldn't get sessions for user with uid %v: %v", userUID, err)
		return nil, err
	}
	splitAuthorization := strings.Split(authorizationString, "|")
	if len(splitAuthorization) != 2 {
		logger.Warn("User with uid %v sent a bad header authorization", userUID)
		return nil, serverErr.ErrBadRequest
	}
	sentTime := splitAuthorization[0]
	hashedTime := splitAuthorization[1]
	t, err := time.Parse(time.RFC3339, sentTime)
	if err != nil {
		logger.Warn("User with uid %v sent a bad header authorization: %v", userUID, err)
	}

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	if t.Before(now.Add(-allowedTimeDifference)) || t.After(now.Add(allowedTimeDifference)) {
		logger.Warn("User with uid %v is using invalid time in header", userUID)
		return nil, serverErr.ErrBadRequest
	}

	for _, session := range sessions {
		hash := md5.New()
		hash.Write([]byte(sentTime))
		hash.Write([]byte(session.Token))
		h := hash.Sum(nil)
		if hex.EncodeToString(h) == hashedTime {
			return session, nil
		}
	}
	logger.Warn("User with uid %v is using bad authentication token", userUID)
	return nil, serverErr.ErrNotAuthenticated
}

func removeSession(ctx context.Context, authorizationString, userUID string) (err error) {
	session, err := getSession(ctx, authorizationString, userUID)
	if err != nil {
		logger.Error("Couldn't get session for user with uid %v: %v", userUID, err)
		return
	}
	err = data.RemoveSession(ctx, session.Token)
	if err != nil {
		logger.Error("Couldn't remove session for user with uid %v: %v", userUID, err)
		return
	}
	return
}

func parseFile(ctx context.Context, filePath string) (parsedQuestions *dto.TestQuestions, err error) {
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		logger.Error("Couldn't open file", err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered in ", r)
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

func _generateToken() string {
	raw := make([]byte, 18, 18)
	enc := make([]byte, 24, 24)

	rand.Read(raw)
	base64.RawURLEncoding.Encode(enc, raw)

	return string(enc)
}
