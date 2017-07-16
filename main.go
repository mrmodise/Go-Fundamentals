package main

import "./greeting"

func main() {
	var s = greeting.Salutation{}
	s.Greeting = "Hi"
	s.Name = "Thabo"
	greeting.Greet(s, greeting.CustomPrint("!"), true)
}
