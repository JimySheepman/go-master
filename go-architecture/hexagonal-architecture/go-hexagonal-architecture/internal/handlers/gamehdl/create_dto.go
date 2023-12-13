package gamehdl

import "go-hexagonal-architecture/internal/core/domain"

type BodyCreate struct {
	Name  string `json:"name"`
	Size  uint   `json:"size"`
	Bombs uint   `json:"bombs"`
}

type ResponseCreate domain.Game

func BuildResponseCreate(model domain.Game) ResponseCreate {
	return ResponseCreate(model)
}
