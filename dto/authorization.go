package dto

import "projekat/enum"

type SessionInfo struct {
	UserUID  string
	Username string
	Role     enum.Role
	Token    string
}
