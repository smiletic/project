package dto

import (
	"time"
)

type CreateExaminationResponse struct {
	UID string `json:"Uid"`
}

type CreateExaminationRequest struct {
	DoctorUID       string    `json:"DoctorUid"`
	PatientUID      string    `json:"PatientUid"`
	ExaminationDate time.Time `json:",string"`
}

type ExaminationInfo struct {
	UID             string    `json:"Uid"`
	DoctorUID       string    `json:"DoctorUid"`
	PatientUID      string    `json:"PatientUid"`
	ExaminationDate time.Time `json:",string"`
}

type GetExaminationsResponse struct {
	Examinations []*ExaminationInfo
}
