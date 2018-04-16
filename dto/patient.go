package dto

import (
	"time"
)

type CreatePatientResponse struct {
	UID string `json:"Uid"`
}

type CreatePatientRequest struct {
	PersonUID            string    `json:"PersonUid"`
	MedicalRecordID      string    `json:"MedicalRecordId"`
	HealthCardID         string    `json:"HealthCardId"`
	HealthCardValidUntil time.Time `json:",string"`
}

type UpdatePatientRequest struct {
	MedicalRecordID      string    `json:"MedicalRecordId"`
	HealthCardID         string    `json:"HealthCardId"`
	HealthCardValidUntil time.Time `json:",string"`
}

type GetPatientResponse struct {
	UID                  string `json:"Uid"`
	PersonUID            string `json:"PersonUid"`
	PersonName           string
	PersonSurname        string
	MedicalRecordID      string    `json:"MedicalRecordId"`
	HealthCardID         string    `json:"HealthCardId"`
	HealthCardValidUntil time.Time `json:",string"`
}

type GetPatientsResponse struct {
	Patients []*GetPatientResponse
}
