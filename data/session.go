package data

import (
	"context"
	"fmt"
	"masterRad/db"
)

var (
	CheckLogin = checkLogin
)

func checkLogin(ctx context.Context, name, passhash string) (userExists bool) {
	d := ctx.Value(db.RunnerKey).(db.Runner)

	query := `
		select *
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

	return rr.ScanNext()
}
