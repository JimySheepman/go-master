//go:generate mockgen -package=mockups -destination=../../../mocks/mockups/repositories.go -source=repositories.go

package ports

import "go-hexagonal-architecture/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
