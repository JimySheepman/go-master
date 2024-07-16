package behavioral

import "fmt"

type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
}

type BaseGame struct {
	Game
}

func (b *BaseGame) Play() {
	b.Initialize()
	b.StartPlay()
	b.EndPlay()
}

type Cricket struct {
	BaseGame
}

func (c *Cricket) Initialize() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}

func (c *Cricket) StartPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (c *Cricket) EndPlay() {
	fmt.Println("Cricket Game Finished!")
}

type Football struct {
	BaseGame
}

func (f *Football) Initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}

func (f *Football) StartPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (f *Football) EndPlay() {
	fmt.Println("Football Game Finished!")
}

func TemplateMethod() {
	cricket := &Cricket{}
	cricket.BaseGame.Game = cricket
	cricket.Play()

	fmt.Println()

	football := &Football{}
	football.BaseGame.Game = football
	football.Play()
}
