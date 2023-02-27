package userModel

import "kbrprime-be/internal/app/model/helperModel"

type UserOTP struct {
	ID         int    `json:"id"`
	UserID     int64  `json:"user_id"`
	OTPCode    string `json:"otp_code"`
	ExpiredAt  int64  `json:"expired_at"`
	ModuleName string `json:"module_name"`
	TypeOTP    string `json:"type_otp"`
	SendTo     string `json:"send_to"`
	SendAt     int64  `json:"send_at"`
	Response   string `json:"response"`
	ProductID  string `json:"product_id"`
	helperModel.UserAuditModel
	helperModel.DateAuditModel
}
