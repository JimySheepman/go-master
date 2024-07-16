package creational

import "fmt"

type House struct {
	doorType   string
	windowType string
	floor      int
}

type HouseBuilder struct {
	doorType   string
	windowType string
	floor      int
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) SetDoorType(doorType string) *HouseBuilder {
	b.doorType = doorType
	return b
}

func (b *HouseBuilder) SetWindowType(windowType string) *HouseBuilder {
	b.windowType = windowType
	return b
}

func (b *HouseBuilder) SetFloor(floor int) *HouseBuilder {
	b.floor = floor
	return b
}

func (b *HouseBuilder) Build() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func Builder() {
	builder := NewHouseBuilder()
	house := builder.SetDoorType("Wooden Door").
		SetWindowType("Sliding Window").
		SetFloor(2).
		Build()

	fmt.Println(house)
}
