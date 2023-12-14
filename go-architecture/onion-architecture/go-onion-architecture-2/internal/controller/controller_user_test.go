package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang_startup/internal/model"
	"golang_startup/internal/model/request"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	idTest    = "00000000-0000-0000-0000-000000000000"
	emailTest = "email@domain"
	nameTest  = "Test"

	jsonOutput = "{\"id\":\"00000000-0000-0000-0000-000000000000\",\"name\":\"Test\"," +
		"\"email\":\"email@domain\",\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"0001-01-01T00:00:00Z\"}\n"
)

var testUser = model.User{
	ID:    idTest,
	Name:  nameTest,
	Email: emailTest,
}

func TestServer_FindUser(t *testing.T) {
	t.Run("No error return", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Equal(t, w.Body.String(), jsonOutput)
	})

	t.Run("Error GetUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			FindUser(gomock.Any(), idTest).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestServer_CreateUser(t *testing.T) {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(testUser)
	if err != nil {
		require.NoError(t, err)
	}

	t.Run("Create user successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			CreateUser(gomock.Any(), testUser).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/user", serverReturn.URL)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPost, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, w.Body.String(), buf.String())
	})

	t.Run("Error creating user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			CreateUser(gomock.Any(), testUser).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/user", serverReturn.URL)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPost, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestServer_UpdateUser(t *testing.T) {
	var (
		buf bytes.Buffer

		updateUser = request.UserUpdate{
			Name:  nameTest,
			Email: emailTest,
		}
	)

	err := json.NewEncoder(&buf).Encode(testUser)
	if err != nil {
		require.NoError(t, err)
	}

	t.Run("Update user successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			UpdateUser(gomock.Any(), idTest, updateUser).
			Return(testUser, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPut, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, w.Body.String(), buf.String())
	})

	t.Run("Error updating user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			UpdateUser(gomock.Any(), idTest, updateUser).
			Return(model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/user/%s", serverReturn.URL, idTest)

		reader := bytes.NewReader(buf.Bytes())

		r, _ := http.NewRequest(http.MethodPut, serverURL, reader)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestServer_FindUsers(t *testing.T) {
	testUsers := []model.User{
		testUser, testUser,
	}

	const (
		testLimit  int64 = 20
		testOffset int64 = 0
	)

	t.Run("Find all users successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		mockService.EXPECT().
			FindAllUsers(gomock.Any(), testOffset, testLimit).
			Return(testUsers, nil).
			Times(1)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/users?offset=%d&limit=%d", serverReturn.URL, testOffset, testLimit)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)

		jsonUsers, err := json.Marshal(testUsers)
		require.NoError(t, err)

		assert.Contains(t, w.Body.String(), string(jsonUsers))
	})

	t.Run("Error fetching all users", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		mockService.EXPECT().
			FindAllUsers(gomock.Any(), testOffset, testLimit).
			Return([]model.User{}, fmt.Errorf("error")).
			Times(1)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/users?offset=%d&limit=%d", serverReturn.URL, testOffset, testLimit)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Error bad request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := NewMockService(ctrl)

		serverTest := New(mockService)

		serverReturn := httptest.NewServer(serverTest.mux)
		serverURL := fmt.Sprintf("%s/users?offset=%d&limit=sasasa", serverReturn.URL, testOffset)

		r, _ := http.NewRequest(http.MethodGet, serverURL, nil)
		w := httptest.NewRecorder()

		serverTest.mux.ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
