package repository

import (
	"github.com/arvinpaundra/clean-architecture/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (u *UserRepositoryMock) Create(user model.User) error {
	ret := u.Mock.Called(user)

	return ret.Error(0)
}

func (u *UserRepositoryMock) Login(email string, password string) (model.User, error) {
	ret := u.Mock.Called(email, password)

	return ret.Get(0).(model.User), ret.Error(1)
}

func (u *UserRepositoryMock) GetAll() ([]model.User, error) {
	ret := u.Mock.Called()

	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}

	return ret.Get(0).([]model.User), ret.Error(1)
}
