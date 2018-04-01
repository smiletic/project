package dto

import "masterRad/enum"

type Authorization struct {
	Korisnik_UID string
	Username     string
	Role         enum.Role
}
