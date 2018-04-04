package data

import (
	"context"
	"masterRad/db"
	"masterRad/dto"
)

var (
	CreateEmployee = createEmployee
	UpdateEmployee = updateEmployee
	DeleteEmployee = deleteEmployee
	GetEmployee    = getEmployee
	GetEmployees   = getEmployees
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
				uid as "UID",
				person_uid as "PersonUID",
				work_document_id as "WorkDocumentID"
				from employee
				where uid = $1`

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

func getEmployees(ctx context.Context) (employees *dto.GetEmployeesResponse, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)
	query := `select 
				uid as "UID",
				person_uid as "PersonUID",
				work_document_id as "WorkDocumentID"
				from employee`

	rows, err := d.Query(ctx, query)
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
