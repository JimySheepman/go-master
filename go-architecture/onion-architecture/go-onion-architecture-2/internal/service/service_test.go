package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang_startup/internal/common"
	"golang_startup/internal/model"
	"golang_startup/internal/model/request"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	idTest    = "00000000-0000-0000-0000-000000000000"
	emailTest = "email@domain"
	nameTest  = "Test"

	testOffset int64 = 0
	testLimit  int64 = 20
)

var (
	userTest = model.User{
		ID:    idTest,
		Name:  nameTest,
		Email: emailTest,
	}

	userTestWithUpdatedTimestamps = model.User{
		ID:        idTest,
		Name:      nameTest,
		Email:     emailTest,
		CreatedAt: timeTest,
		UpdatedAt: timeTest,
	}

	timeTest = time.Now()
)

func TestService_FindUser(t *testing.T) {
	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timerMock := NewMockTimer(ctrl)

		service := New(repositoryMock, timerMock)

		repositoryMock.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		_, err := service.FindUser(context.Background(), idTest)
		assert.Error(t, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timerMock := NewMockTimer(ctrl)

		service := New(repositoryMock, timerMock)

		repositoryMock.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(userTest, nil).
			Times(1)

		user, err := service.FindUser(context.Background(), idTest)
		assert.NoError(t, err)
		assert.Equal(t, userTest, user)
	})
}

func TestService_CreateUser(t *testing.T) {
	t.Run("Returns an error ", func(t *testing.T) {
		t.Run("in create user method.", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := NewMockRepository(ctrl)
			timerMock := NewMockTimer(ctrl)

			service := New(repositoryMock, timerMock)

			gomock.InOrder(
				timerMock.EXPECT().
					Now().
					Return(timeTest).
					Times(2),
				repositoryMock.EXPECT().
					CreateUser(gomock.Any(), userTestWithUpdatedTimestamps).
					Return(idTest, fmt.Errorf("error")).
					Times(1),
			)

			_, err := service.CreateUser(context.Background(), userTest)
			assert.Error(t, err)
		})

		t.Run("in find user method", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := NewMockRepository(ctrl)
			timerMock := NewMockTimer(ctrl)

			service := New(repositoryMock, timerMock)

			gomock.InOrder(
				timerMock.EXPECT().
					Now().
					Return(timeTest).
					Times(2),
				repositoryMock.EXPECT().
					CreateUser(gomock.Any(), userTestWithUpdatedTimestamps).
					Return(idTest, nil).
					Times(1),
				repositoryMock.EXPECT().
					FindUser(gomock.Any(), idTest).
					Return(model.User{}, fmt.Errorf("error")).
					Times(1),
			)

			_, err := service.CreateUser(context.Background(), userTest)
			assert.Error(t, err)
		})
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timerMock := NewMockTimer(ctrl)

		service := New(repositoryMock, timerMock)

		gomock.InOrder(
			timerMock.EXPECT().
				Now().
				Return(timeTest).
				Times(2),
			repositoryMock.EXPECT().
				CreateUser(gomock.Any(), userTestWithUpdatedTimestamps).
				Return(idTest, nil).
				Times(1),
			repositoryMock.EXPECT().
				FindUser(gomock.Any(), idTest).
				Return(userTest, nil).
				Times(1),
		)

		user, err := service.CreateUser(context.Background(), userTest)
		assert.NoError(t, err)
		assert.Equal(t, userTest, user)
	})
}

func TestService_UpdateUser(t *testing.T) {
	const (
		emailUpdate = "email update"
		nameUpdate  = "name update"
	)

	var (
		userUpdateRequest = request.UserUpdate{
			Name:  nameUpdate,
			Email: emailUpdate,
		}
		expectedUserUpdated = model.User{
			ID:        idTest,
			Name:      nameUpdate,
			Email:     emailUpdate,
			UpdatedAt: timeTest,
		}
	)

	t.Run("Returns an error", func(t *testing.T) {
		t.Run("in FindUser repo", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := NewMockRepository(ctrl)
			timerMock := NewMockTimer(ctrl)

			service := New(repositoryMock, timerMock)

			repositoryMock.EXPECT().
				FindUser(gomock.Any(), idTest).
				Return(model.User{}, fmt.Errorf("error")).
				Times(1)

			user, err := service.UpdateUser(context.Background(), idTest, userUpdateRequest)
			assert.Error(t, err)

			assert.Equal(t, model.User{}, user)
		})
		t.Run("in UpdateUser repo", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryMock := NewMockRepository(ctrl)
			timerMock := NewMockTimer(ctrl)

			service := New(repositoryMock, timerMock)

			gomock.InOrder(
				repositoryMock.EXPECT().
					FindUser(gomock.Any(), idTest).
					Return(userTest, nil).
					Times(1),
				timerMock.EXPECT().
					Now().
					Return(timeTest).
					Times(1),
				repositoryMock.EXPECT().
					UpdateUser(gomock.Any(), expectedUserUpdated).
					Return(fmt.Errorf("error")).
					Times(1),
			)

			user, err := service.UpdateUser(context.Background(), idTest, userUpdateRequest)
			assert.Error(t, err)

			assert.Equal(t, model.User{}, user)
		})
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timerMock := NewMockTimer(ctrl)
		service := New(repositoryMock, timerMock)

		gomock.InOrder(
			repositoryMock.EXPECT().
				FindUser(gomock.Any(), idTest).
				Return(userTest, nil).
				Times(1),
			timerMock.EXPECT().
				Now().
				Return(timeTest).
				Times(1),
			repositoryMock.EXPECT().
				UpdateUser(gomock.Any(), expectedUserUpdated).
				Return(nil).
				Times(1),
		)

		user, err := service.UpdateUser(context.Background(), idTest, userUpdateRequest)
		assert.NoError(t, err)
		assert.Equal(t, expectedUserUpdated, user)
	})
}

func TestService_DeleteUser(t *testing.T) {
	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timer := common.NewTimer()

		service := New(repositoryMock, timer)

		repositoryMock.EXPECT().
			DeleteUser(gomock.Any(), idTest).
			Return(fmt.Errorf("error")).
			Times(1)

		err := service.DeleteUser(context.Background(), idTest)
		assert.Error(t, err)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timerMock := NewMockTimer(ctrl)

		service := New(repositoryMock, timerMock)

		repositoryMock.EXPECT().
			DeleteUser(gomock.Any(), idTest).
			Return(nil).
			Times(1)

		err := service.DeleteUser(context.Background(), idTest)
		assert.NoError(t, err)
	})
}

func TestService_FindAllUsers(t *testing.T) {
	t.Run("Returns an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repositoryMock := NewMockRepository(ctrl)
		timer := common.NewTimer()

		service := New(repositoryMock, timer)

		repositoryMock.EXPECT().
			FindAllUsers(gomock.Any(), testOffset, testLimit).
			Return([]model.User{}, fmt.Errorf("error")).
			Times(1)

		users, err := service.FindAllUsers(context.Background(), testOffset, testLimit)
		assert.Error(t, err)
		assert.Empty(t, users)
	})

	t.Run("Returns no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		testUsers := []model.User{userTest, userTest}

		repositoryMock := NewMockRepository(ctrl)
		timerMock := NewMockTimer(ctrl)

		service := New(repositoryMock, timerMock)

		repositoryMock.EXPECT().
			FindAllUsers(gomock.Any(), testOffset, testLimit).
			Return(testUsers, nil).
			Times(1)

		users, err := service.FindAllUsers(context.Background(), testOffset, testLimit)
		assert.NoError(t, err)
		assert.Len(t, users, 2)
	})
}
