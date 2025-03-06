package structural

import "fmt"

type Employee interface {
	ShowDetails()
}

type Developer struct {
	name string
}

func (d *Developer) ShowDetails() {
	fmt.Println("Developer:", d.name)
}

type Manager struct {
	name string
}

func (m *Manager) ShowDetails() {
	fmt.Println("Manager:", m.name)
}

type Company struct {
	employees []Employee
}

func (c *Company) AddEmployee(e Employee) {
	c.employees = append(c.employees, e)
}

func (c *Company) ShowDetails() {
	for _, e := range c.employees {
		e.ShowDetails()
	}
}

func Composite() {
	dev1 := &Developer{name: "Alice"}
	dev2 := &Developer{name: "Bob"}
	manager := &Manager{name: "Charlie"}

	company := &Company{}
	company.AddEmployee(dev1)
	company.AddEmployee(dev2)
	company.AddEmployee(manager)

	company.ShowDetails()
}
