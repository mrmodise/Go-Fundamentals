package greeting

import ("fmt")

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

const (
	PI = 3.14
	LANGUAGE = "Go"
	A = iota
	B = iota
	C = iota
)

func CustomPrint(s string) Printer {
	return func(str string) {
		fmt.Println(str, s)
	}
}

func GetPrefix(name string) (prefix string)  {
	switch name {
	case "Modise": prefix = "Mr "
	case "Nwabisa": prefix = "Mrs "
	case "Thuso": prefix = "Mr "
	case "Thabo": prefix = "Miss "
	default: prefix = "Dude "
	}
	return
}

func Greet(salute Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salute.Name, salute.Greeting)

	if prefix := GetPrefix(salute.Name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hello " + name
	return
}
