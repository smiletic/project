package data

import (
	"context"
	"masterRad/db"
	"masterRad/dto"
)

var (
	CreatePatient              = createPatient
	UpdatePatient              = updatePatient
	DeletePatient              = deletePatient
	GetPatient                 = getPatient
	GetPatientsByName          = getPatientsByName
	GetPatientsByHealthDocUIDs = getPatientsByHealthDocUIDs
)

func createPatient(ctx context.Context, request *dto.CreatePatientRequest) (uid string, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `insert into patient 
				(person_uid, medical_record_id, health_card_id, health_card_valid_until) 
				values ($1, $2, $3, $4)
				returning uid`

	rows, err := d.Query(ctx, query, request.PersonUID, request.MedicalRecordID, request.HealthCardID, request.HealthCardValidUntil)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		uid = rr.ReadByIdxString(0)
	}

	err = rr.Error()
	return
}

func updatePatient(ctx context.Context, patientUID string, request *dto.UpdatePatientRequest) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `update patient set
				medical_record_id = $1, health_card_id = $2, health_card_valid_until = $3
				where uid = $4`

	_, err = d.Exec(ctx, query, request.MedicalRecordID, request.HealthCardID, request.HealthCardValidUntil, patientUID)

	return
}

func deletePatient(ctx context.Context, patientUID string) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `delete from patient
				where uid = $1`

	_, err = d.Exec(ctx, query, patientUID)

	return
}

func getPatient(ctx context.Context, patientUID string) (patient *dto.GetPatientResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `select 
				p.uid as "UID",
				p.person_uid as "PersonUID",
				pe.name as "PersonName",
				pe.surname as "PersonSurname",
				p.medical_record_id as "MedicalRecordID",
				p.health_card_id as "HealthCardID",
				p.health_card_valid_until as "HealthCardValidUntil"
				from patient p
				join person pe on (p.person_uid = pe.uid) 
				where p.uid = $1`

	rows, err := d.Query(ctx, query, patientUID)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		patient = &dto.GetPatientResponse{}
		rr.ReadAllToStruct(patient)
	}

	err = rr.Error()
	return
}

func getPatientsByName(ctx context.Context, name, surname string) (patients *dto.GetPatientsResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)
	name = "%" + name + "%"
	surname = "%" + surname + "%"
	query := `select 
				p.uid as "UID",
				p.person_uid as "PersonUID",
				pe.name as "PersonName",
				pe.surname as "PersonSurname",
				p.medical_record_id as "MedicalRecordID",
				p.health_card_id as "HealthCardID",
				p.health_card_valid_until as "HealthCardValidUntil"
				from patient p 
				join person pe on (p.person_uid = pe.uid) 
				where pe.name ilike $1
				and pe.surname ilike $2`

	rows, err := d.Query(ctx, query, name, surname)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}
	patients = &dto.GetPatientsResponse{}
	for rr.ScanNext() {
		patient := &dto.GetPatientResponse{}
		rr.ReadAllToStruct(patient)
		patients.Patients = append(patients.Patients, patient)
	}
	err = rr.Error()
	return
}

func getPatientsByHealthDocUIDs(ctx context.Context, medicalRecordID, healthCardID string) (patients *dto.GetPatientsResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)
	medicalRecordID = "%" + medicalRecordID + "%"
	healthCardID = "%" + healthCardID + "%"
	query := `select 
				p.uid as "UID",
				p.person_uid as "PersonUID",
				pe.name as "PersonName",
				pe.surname as "PersonSurname",
				p.medical_record_id as "MedicalRecordID",
				p.health_card_id as "HealthCardID",
				p.health_card_valid_until as "HealthCardValidUntil"
				from patient p 
				join person pe on (p.person_uid = pe.uid) 
				where p.medical_record_id ilike $1
				and p.health_card_id ilike $2`

	rows, err := d.Query(ctx, query, medicalRecordID, healthCardID)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}
	patients = &dto.GetPatientsResponse{}
	for rr.ScanNext() {
		patient := &dto.GetPatientResponse{}
		rr.ReadAllToStruct(patient)
		patients.Patients = append(patients.Patients, patient)
	}
	err = rr.Error()
	return
}
