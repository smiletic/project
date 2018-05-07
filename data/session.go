package data

import (
	"context"
	"fmt"
	"masterRad/db"
	"masterRad/dto"
	"masterRad/enum"
)

var (
	Login = login
)

func login(ctx context.Context, name, passhash string) (autorizacija *dto.Authorization, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `
		select uid, role, username
		from system_user
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
