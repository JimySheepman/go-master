package structural

import "fmt"

type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU started")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory loaded")
}

type HardDrive struct{}

func (hd *HardDrive) Read() {
	fmt.Println("Hard drive reading data")
}

type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputer() *Computer {
	return &Computer{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *Computer) Start() {
	c.cpu.Start()
	c.memory.Load()
	c.hardDrive.Read()
}

func Facade() {
	computer := NewComputer()
	computer.Start()
}
