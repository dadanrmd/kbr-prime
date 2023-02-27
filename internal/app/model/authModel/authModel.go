package authModel

import (
	"kbrprime-be/internal/app/model/helperModel"
	"kbrprime-be/internal/app/model/userModel"
)

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	User  UserLogin `json:"user"`
	Token string    `json:"token"`
}

type UserLogin struct {
	IdUser     int64  `json:"id_user"`
	Nik        string `json:"nik"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	NoHp       string `json:"no_hp"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	Update_Key string `json:"update_key"`
}

type ChangePasswordReq struct {
	User        userModel.User `json:"-"`
	OldPassword string         `json:"old_password" binding:"required"`
	NewPassword string         `json:"new_password" binding:"required"`
}

type ChangePasswordRes struct {
}

type ForgotPasswordReq struct {
	Email string `json:"email" binding:"required"`
}

type ForgotPasswordRes struct {
}

type VerifyForgotPasswordReq struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type VerifyForgotPasswordRes struct {
}

type GetUserFromJWTAndRoleReq struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type GetUserFromJWTAndRoleRes struct {
	User userModel.User
}

type TokenValidityReq struct {
	Token string `json:"token"`
}

type TokenValidityRes struct {
	IsValid bool `json:"is_valid"`
}

type ChangePasswordFromForgotPassReq struct {
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ChangePasswordFromForgotPassRes struct {
}

type TokenSession struct {
	ID        int64  `json:"id"`
	TokenUID  string `json:"token_uid"`
	IsRevoked bool   `json:"is_revoked"`
	helperModel.DateAuditModel
	helperModel.UserAuditModel
}
