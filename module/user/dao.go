package user

import (
	"davinci/client"
	"fmt"
)

type IUserDao interface {
	FindUserAll() []*Model
	FindUserById(id int64, user *Model) (*Model, error)
	FindUserByUsername(username string) (*Model, error)
	FindUserByEmail(email string) (*Model, error)
	CreateUser(user *Model) error
	UpdateUser(id int, user *Model) error
	DeleteUser(id int) error
}

var DaoGetter IUserDao

func init() {
	DaoGetter = NewGetterImpl()
}

type GetterImpl struct {
}

func NewGetterImpl() *GetterImpl {
	return &GetterImpl{}
}

func (I GetterImpl) FindUserByUsername(username string) (*Model, error) {
	var user Model
	err := client.Orm.Where("username=?", username).Find(&user).Error

	return &user, err
}

func (I GetterImpl) FindUserByEmail(email string) (*Model, error) {
	var user Model
	err := client.Orm.Where("email=?", email).Find(&user).Error
	return &user, err
}

func (I GetterImpl) CreateUser(user *Model) error {
	return client.Orm.Create(user).Error
}

func (I GetterImpl) UpdateUser(id int, user *Model) error {
	//TODO implement me
	panic("implement me")
}

func (I GetterImpl) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (I GetterImpl) FindUserById(id int64, user *Model) (*Model, error) {
	//TODO implement me
	db := client.Orm.Where("id=?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found, id=%d", id)
	}
	return user, nil
}

func (I GetterImpl) FindUserAll() []*Model {
	var users []*Model
	client.Orm.Find(&users)
	return users
}
