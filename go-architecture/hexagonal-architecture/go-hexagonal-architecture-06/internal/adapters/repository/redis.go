package repository

import (
	"Messenger/internal/core/domain"
	"encoding/json"

	"github.com/go-redis/redis/v7"
)

type MessengerRedisRepository struct {
	client *redis.Client
}

func NewMessengerRedisRepository(host string) *MessengerRedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	return &MessengerRedisRepository{
		client: client,
	}
}

func (r *MessengerRedisRepository) SaveMessage(message domain.Message) error {
	json, err := json.Marshal(message)
	if err != nil {
		return err
	}

	r.client.HSet("messages", message.ID, json)

	return nil
}

func (r *MessengerRedisRepository) ReadMessage(id string) (*domain.Message, error) {
	value, err := r.client.HGet("messages", id).Result()
	if err != nil {
		return nil, err
	}

	message := &domain.Message{}

	if err := json.Unmarshal([]byte(value), message); err != nil {
		return nil, err
	}

	return message, nil
}

func (r *MessengerRedisRepository) ReadMessages() ([]*domain.Message, error) {
	messages := []*domain.Message{}
	value, err := r.client.HGetAll("messages").Result()
	if err != nil {
		return nil, err
	}

	for _, val := range value {
		message := &domain.Message{}
		if err = json.Unmarshal([]byte(val), message); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}
