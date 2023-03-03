package auth

import "project-intern-bcc/src/business/entity"

type UserAuthInfo struct {
	User  entity.UserResponse
	Token string
}