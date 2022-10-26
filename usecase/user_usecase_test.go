package usecase

import (
	"github.com/arvinpaundra/clean-architecture/dto"
	"github.com/arvinpaundra/clean-architecture/model"
	"github.com/arvinpaundra/clean-architecture/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var userRepository = repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUsecaseTest = NewUserUsecase(&userRepository)

func TestUserUsecase_GetAll(t *testing.T) {
	users := []model.User{
		{
			Email:    "arvin@mail.com",
			Password: "123",
		},
	}

	userRepository.Mock.On("GetAll").Return(users, nil)

	result, err := userUsecaseTest.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, users[0].Email, result[0].Email)
	assert.Equal(t, users[0].Password, result[0].Password)
}

func TestUserUsecase_Create(t *testing.T) {
	user := model.User{
		Email:    "arvin@mail.com",
		Password: "123",
	}

	userDTO := dto.UserDTO{
		Email:    "arvin@mail.com",
		Password: "123",
	}

	userRepository.Mock.On("Create", user).Return(nil)

	err := userUsecaseTest.Create(userDTO)

	assert.Nil(t, err)
}

func TestUserUsecase_Login(t *testing.T) {
	user := model.User{
		Email:    "arvin@mail.com",
		Password: "123",
	}

	userDTO := dto.UserDTO{
		Email:    "arvin@mail.com",
		Password: "123",
	}

	userRepository.Mock.On("Login", "arvin@mail.com", "123").Return(user, nil)

	result, err := userUsecaseTest.Login(userDTO)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Password, result.Password)
}
