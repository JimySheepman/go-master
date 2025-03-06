package solids

import "fmt"

type solid func()

var solids = []solid{
	SingleResponsibilityPrinciple,
	OpenClosedPrinciple,
	LiskovSubstitutionPrinciple,
	InterfaceSegregationPrinciple,
	DependencyInversionPrinciple,
}

var solidsName = map[int]string{
	0: "Single Responsibility Principle",
	1: "Open/Closed Principle",
	2: "Liskov Substitution Principle",
	3: "Interface Segregation Principle",
	4: "Dependency Inversion Principle",
}

type User struct {
	ID   int
	Name string
}

type UserService struct{}

func (us *UserService) GetUser(ID int) *User {
	return &User{ID: ID, Name: "John Doe"}
}

type UserPrinter struct{}

func (up *UserPrinter) PrintUser(user *User) {
	fmt.Printf("User: %d, Name: %s\n", user.ID, user.Name)
}

func SingleResponsibilityPrinciple() {
	us := &UserService{}
	up := &UserPrinter{}

	user := us.GetUser(1)
	up.PrintUser(user)
}

type PaymentProcessor interface {
	Pay(amount float64)
}

type CreditCard struct{}

func (cc *CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using credit card\n", amount)
}

type PayPal struct{}

func (pp *PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal\n", amount)
}

func OpenClosedPrinciple() {
	var processor PaymentProcessor

	processor = &CreditCard{}
	processor.Pay(100.0)

	processor = &PayPal{}
	processor.Pay(150.0)
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

func LiskovSubstitutionPrinciple() {
	animals := []Animal{
		&Dog{},
		&Cat{},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

type Drawable interface {
	Draw()
}

type Printable interface {
	Print()
}

type Graphic struct{}

func (g *Graphic) Draw() {
	fmt.Println("Drawing graphic")
}

func (g *Graphic) Print() {
	fmt.Println("Printing graphic")
}

func InterfaceSegregationPrinciple() {
	var d Drawable = &Graphic{}
	var p Printable = &Graphic{}

	d.Draw()
	p.Print()
}

type MessageSender interface {
	SendMessage(msg string)
}

type EmailSender struct{}

func (es *EmailSender) SendMessage(msg string) {
	fmt.Printf("Sending email: %s\n", msg)
}

type SMSSender struct{}

func (ss *SMSSender) SendMessage(msg string) {
	fmt.Printf("Sending SMS: %s\n", msg)
}

type Notifier struct {
	sender MessageSender
}

func (n *Notifier) Notify(msg string) {
	n.sender.SendMessage(msg)
}

func DependencyInversionPrinciple() {
	emailSender := &EmailSender{}
	smsSender := &SMSSender{}

	emailNotifier := &Notifier{sender: emailSender}
	smsNotifier := &Notifier{sender: smsSender}

	emailNotifier.Notify("Hello via Email")
	smsNotifier.Notify("Hello via SMS")
}

func PrintSolid() {
	for i, sFunc := range solids {
		fmt.Println("Algorithm name:", solidsName[i])
		sFunc()
	}
}
