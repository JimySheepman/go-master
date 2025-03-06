package structural

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	color string
}

func (c *Circle) Draw() {
	fmt.Println("Drawing circle of color", c.color)
}

type ShapeFactory struct {
	circleMap map[string]*Circle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{circleMap: make(map[string]*Circle)}
}

func (sf *ShapeFactory) GetCircle(color string) *Circle {
	if circle, ok := sf.circleMap[color]; ok {
		return circle
	}
	circle := &Circle{color: color}
	sf.circleMap[color] = circle
	return circle
}

func Flyweight() {
	factory := NewShapeFactory()

	redCircle := factory.GetCircle("red")
	redCircle.Draw()

	blueCircle := factory.GetCircle("blue")
	blueCircle.Draw()

	anotherRedCircle := factory.GetCircle("red")
	anotherRedCircle.Draw()

	fmt.Println("Number of unique circles:", len(factory.circleMap)) // 2
}
