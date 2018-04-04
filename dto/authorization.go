package dto

import "masterRad/enum"

type Authorization struct {
	User_UID string
	Username string
	Role     enum.Role
}
