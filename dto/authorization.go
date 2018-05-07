package dto

import "masterRad/enum"

type Authorization struct {
	UserUID  string
	Username string
	Role     enum.Role
}
