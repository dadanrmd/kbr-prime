package userModel

/* Table Definition */
type User struct {
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

func (User) TableName() string {
	return "kbr_users"
}

/* DTO */
type CreateUserReq struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NoHp     string `json:"no_hp"`
}
type UpdateUserReq struct {
	IdUser int64  `json:"id_user"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	NoHp   string `json:"no_hp"`
	Status string `json:"status"`
}
type DeleteUserReq struct {
	ID int64 `json:"id"`
}

const (
	MsgSuccessUpdateUser = "user berhasil diperbaharui"
	MsgSuccessDeleteUser = "user berhasil dihapus"
	MsgFailedAddUser     = "user gagal di tambahkan"
)
