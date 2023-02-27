package userRepository

import (
	"kbrprime-be/internal/app/model/userModel"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{conn: db}
}

func (d userRepository) FindUserByID(id int64) (*userModel.User, error) {
	var userData userModel.User
	db := d.conn.Preload(clause.Associations)
	db.First(&userData, "id=?", id)
	return &userData, db.Error
}

func (d userRepository) FindUserByEmail(email string) (userData *userModel.User, err error) {
	err = d.conn.First(&userData, "role = 'USER' and email = ?", email).Error
	return
}

func (d userRepository) InsertUser(userData userModel.User) (*userModel.User, error) {
	db := d.conn.Create(&userData)
	return &userData, db.Error
}

func (d userRepository) UpdateUser(userData userModel.User) (*userModel.User, error) {
	db := d.conn.Save(&userData)
	return &userData, db.Error
}

func (d userRepository) FindAllUser() (*[]userModel.User, error) {
	var userData []userModel.User
	db := d.conn.Where("role = 'USER'").Find(&userData)
	return &userData, db.Error
}

func (d userRepository) FindUserByPhone(phone string) (userData *userModel.User, err error) {
	err = d.conn.First(&userData, "role = 'USER' and no_hp = ?", phone).Error
	return
}
