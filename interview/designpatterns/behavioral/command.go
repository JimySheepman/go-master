package behavioral

import "fmt"

type Commander interface {
	Execute()
}

type LightReceiver struct{}

func (l *LightReceiver) On() {
	fmt.Println("Light is ON")
}

func (l *LightReceiver) Off() {
	fmt.Println("Light is OFF")
}

type LightOnCommand struct {
	light *LightReceiver
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *LightReceiver
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type RemoteControlInvoker struct {
	command Commander
}

func (r *RemoteControlInvoker) SetCommand(command Commander) {
	r.command = command
}

func (r *RemoteControlInvoker) PressButton() {
	r.command.Execute()
}

func Command() {
	light := &LightReceiver{}
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	remote := &RemoteControlInvoker{}
	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()
}
