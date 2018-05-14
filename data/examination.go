package data

import (
	"context"
	"masterRad/db"
	"masterRad/dto"
)

var (
	CreateExamination = createExamination
	DeleteExamination = deleteExamination
	GetExaminations   = getExaminations
)

func createExamination(ctx context.Context, request *dto.CreateExaminationRequest) (uid string, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `insert into examination 
				(doctor_uid, patient_uid, date_of_examination) 
				values ($1, $2, $3)
				returning uid`

	rows, err := d.Query(ctx, query, request.DoctorUID, request.PatientUID, request.ExaminationDate)
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

func deleteExamination(ctx context.Context, examinationUID string) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `delete from examination
				where uid = $1`

	_, err = d.Exec(ctx, query, examinationUID)

	return
}

func getExaminations(ctx context.Context) (examinations *dto.GetExaminationsResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)
	query := `	select 
				uid as "UID",
				doctor_uid as "DoctorUID",
				patient_uid as "PatientUID",
				examination_date as "ExaminationDate"
	 			from examination `

	rows, err := d.Query(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}
	examinationInfos := make([]*dto.ExaminationInfo, 0)
	for rr.ScanNext() {
		examination := &dto.ExaminationInfo{}
		rr.ReadAllToStruct(examination)
		examinationInfos = append(examinationInfos, examination)
	}
	examinations.Examinations = examinationInfos
	err = rr.Error()
	return
}
