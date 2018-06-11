package dto

import (
	"encoding/json"
	"time"
)

type CreateFilledRequest struct {
	TestUID        string `json:"TestUid"`
	ExaminationUID string `json:"ExaminationUID"`
	Answers        json.RawMessage
}

type GetFilledTestResponse struct {
	TestUID        string `json:"TestUid"`
	ExaminationUID string `json:"ExaminationUID"`
	Answers        json.RawMessage
}

type GetFilledTestsResponse struct {
	FilledTests []*FilledTestsInfo
}

type FilledTestsInfo struct {
	UID             string `json:"Uid"`
	TestName        string
	ExaminationDate time.Time `json:",string"`
	PatientName     string
	PatientUID      string `json:"PatientUid"`
}
