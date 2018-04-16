package data

import (
	"context"
	"masterRad/db"
	"masterRad/dto"
)

var (
	CreateEmployee               = createEmployee
	UpdateEmployee               = updateEmployee
	DeleteEmployee               = deleteEmployee
	GetEmployee                  = getEmployee
	GetEmployeesByName           = getEmployeesByName
	GetEmployeesByWorkDocumentID = getEmployeesByWorkDocumentID
)

func createEmployee(ctx context.Context, request *dto.CreateEmployeeRequest) (uid string, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `insert into employee 
				(person_uid, work_document_id) 
				values ($1, $2)
				returning uid`

	rows, err := d.Query(ctx, query, request.PersonUID, request.WorkDocumentID)
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

func updateEmployee(ctx context.Context, employeeUID string, request *dto.UpdateEmployeeRequest) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `update employee set
				work_document_id = $1
				where uid = $2`

	_, err = d.Exec(ctx, query, request.WorkDocumentID, employeeUID)

	return
}

func deleteEmployee(ctx context.Context, employeeUID string) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `delete from employee
				where uid = $1`

	_, err = d.Exec(ctx, query, employeeUID)

	return
}

func getEmployee(ctx context.Context, employeeUID string) (employee *dto.GetEmployeeResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `select 
				e.uid as "UID",
				e.person_uid as "PersonUID",
				pe.name as "PersonName",
				pe.surname as "PersonSurname",
				e.work_document_id as "WorkDocumentID"
				from employee e
				join person pe on (e.person_uid = pe.uid) 
				where e.uid = $1`

	rows, err := d.Query(ctx, query, employeeUID)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		employee = &dto.GetEmployeeResponse{}
		rr.ReadAllToStruct(employee)
	}

	err = rr.Error()
	return
}

func getEmployeesByName(ctx context.Context, name, surname string) (employees *dto.GetEmployeesResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)
	name = "%" + name + "%"
	surname = "%" + surname + "%"
	query := `select 
				e.uid as "UID",
				e.person_uid as "PersonUID",
				pe.name as "PersonName",
				pe.surname as "PersonSurname",
				e.work_document_id as "WorkDocumentID"
				from employee e
				join person pe on (e.person_uid = pe.uid)
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
	employees = &dto.GetEmployeesResponse{}
	for rr.ScanNext() {
		employee := &dto.GetEmployeeResponse{}
		rr.ReadAllToStruct(employee)
		employees.Employees = append(employees.Employees, employee)
	}
	err = rr.Error()
	return
}

func getEmployeesByWorkDocumentID(ctx context.Context, workDocID string) (employees *dto.GetEmployeesResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)
	workDocID = "%" + workDocID + "%"

	query := `select 
				e.uid as "UID",
				e.person_uid as "PersonUID",
				pe.name as "PersonName",
				pe.surname as "PersonSurname",
				e.work_document_id as "WorkDocumentID"
				from employee e
				join person pe on (e.person_uid = pe.uid)
				where e.work_document_id ilike $1`

	rows, err := d.Query(ctx, query, workDocID)
	if err != nil {
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}
	employees = &dto.GetEmployeesResponse{}
	for rr.ScanNext() {
		employee := &dto.GetEmployeeResponse{}
		rr.ReadAllToStruct(employee)
		employees.Employees = append(employees.Employees, employee)
	}
	err = rr.Error()
	return
}
