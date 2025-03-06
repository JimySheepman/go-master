package gamesrepo

import (
	"encoding/json"
	"go-hexagonal-architecture/internal/core/domain"
	"go-hexagonal-architecture/pkg/apperrors"

	"github.com/matiasvarela/errors"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{
		kvs: map[string][]byte{},
	}
}

func (repo *memkvs) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {

		var game = domain.Game{}
		if err := json.Unmarshal(value, &game); err != nil {
			return domain.Game{}, errors.New(apperrors.Internal, err, "fail to get value from kvs", "")
		}

		return game, nil
	}

	return domain.Game{}, errors.New(apperrors.NotFound, nil, "game not found in kvs", "")
}

func (repo *memkvs) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New(apperrors.InvalidInput, err, "game fails at marshal into json string", "")
	}

	repo.kvs[game.ID] = bytes

	return nil
}
