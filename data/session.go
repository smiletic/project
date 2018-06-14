package data

import (
	"context"
	"fmt"
	"projekat/db"
	"projekat/dto"
	"projekat/enum"
)

var (
	Login         = login
	CreateSession = createSession
	GetSession    = getSession
	RemoveSession = removeSession
)

func login(ctx context.Context, name, passhash string) (autorizacija *dto.Authorization, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `
		select s.uid, e.role_id, s.username
		from system_user s
		join employee e on (s.employee_uid = e.uid)
		where username = $1
		and password = $2`

	rows, err := d.Query(ctx, query, name, passhash)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		autorizacija = &dto.Authorization{}

		autorizacija.UserUID = rr.ReadByIdxString(0)
		autorizacija.Role = enum.Role(rr.ReadByIdxInt64(1))
		autorizacija.Username = rr.ReadByIdxString(2)
	}

	err = rr.Error()
	return
}

func createSession(ctx context.Context, userUID, token string) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `insert into login_session 
				(system_user_uid, token) 
				values ($1, $2)`

	_, err = d.Exec(ctx, query, userUID, token)

	return
}

func removeSession(ctx context.Context, token string) (err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `delete from login_session 
			  where token = $1`

	_, err = d.Exec(ctx, query, token)

	return
}

func getSession(ctx context.Context, authorization string) (autorizacija *dto.Authorization, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `
		select su.uid, e.role_id, su.username
		from system_user su
		join employee e on (su.employee_uid = e.uid)
		join login_session ls on (su.uid = ls.system_user_uid)
		where ls.token = $1`

	rows, err := d.Query(ctx, query, authorization)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	rr, err := db.GetRowReader(rows)
	if err != nil {
		return
	}

	if rr.ScanNext() {
		autorizacija = &dto.Authorization{}

		autorizacija.UserUID = rr.ReadByIdxString(0)
		autorizacija.Role = enum.Role(rr.ReadByIdxInt64(1))
		autorizacija.Username = rr.ReadByIdxString(2)
	}

	err = rr.Error()
	return
}
