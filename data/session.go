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

func login(ctx context.Context, name, passhash string) (autorizacija *dto.Autorizacija, err error) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `
		select uid, rola, username
		from korisnik
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
		autorizacija = &dto.Autorizacija{}

		autorizacija.Korisnik_UID = rr.ReadByIdxString(0)
		autorizacija.Rola = enum.Rola(rr.ReadByIdxInt64(1))
		autorizacija.Username = rr.ReadByIdxString(2)
	}

	err = rr.Error()
	return
}
