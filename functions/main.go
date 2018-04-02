package main

import (
	"github.com/musale/going-somewhere/functions/closures"
	"github.com/musale/going-somewhere/functions/salut"
)

func main() {
	s := salut.Salute{Name: "Mr King", Greeting: "Ni hao"}
	salut.Greet(s)
	user := closures.User{"King Kunta", "901823"}
	closures.PrintDetails(user, closures.CreatePrinterFunc("GLD"))
}
