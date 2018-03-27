package dto

import "masterRad/enum"

type Autorizacija struct {
	Korisnik_UID string
	Username     string
	Rola         enum.Rola
}
