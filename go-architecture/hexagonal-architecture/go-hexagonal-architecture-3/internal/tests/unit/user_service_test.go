package unit

import (
	"Hexagonal-Architecture/internal/adapters/cache"
	"Hexagonal-Architecture/internal/adapters/repository"
	"Hexagonal-Architecture/internal/core/domain"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setUpDB() *repository.DB {
	db, _ := gorm.Open("postgres", "postgres://test:test@localhost:5433/template1?sslmode=disable")
	db.AutoMigrate(&domain.Message{}, &domain.User{}, &domain.Payment{})

	redisCache, err := cache.NewRedisCache("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	store := repository.NewDB(db, redisCache)

	return store
}

func TestCreateUser(t *testing.T) {
	db := setUpDB()

	email := "alanmoore@example.com"
	password := "password"

	user, err := db.CreateUser(email, password)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}
