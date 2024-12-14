package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() (a string)
	Greet(name string) string
}

type Italian struct{}

type Portuguese struct{}

func (it Italian) LanguageName() (a string) {
	return "Italian"
}

func (it Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

func (por Portuguese) LanguageName() (a string) {
	return "Portuguese"
}

func (por Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %v: %v", g.LanguageName(), g.Greet(name))
}
