package dto

import "projekat/enum"

type Authorization struct {
	UserUID  string
	Username string
	Role     enum.Role
}
