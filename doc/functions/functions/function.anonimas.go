package main

import "fmt"

func main() {
	fmt.Println("Render Tabla de Multiplicar") // functions

	tableof := 9
	renderTable := table(tableof)

	for index := 1; index <= 10; index++ {
		renderTable()
	}
}

func table(value int) func() {

	number := value
	sequence := 0

	return func() {

		sequence += 1
		result := number * sequence
		fmt.Println(sequence, " * ", number, "=", result)
		// return resultado

	}
}

/*
var LoggedEmail string = "david@gmail.com"
var LoggedPass float64 = 134.9
var login func(string, float64) bool = func(email string, password float64) bool {

	if email == LoggedEmail && LoggedPass == password {
		fmt.Println("logeado como admin")
		return true

	}
	fmt.Println("loggen incorrect")
	return false

}

func main() {

	var logged = login("david@gmaifasdfl.com", 134.9)
	if logged {

		login = func(email string, password float64) bool {

			if email != LoggedEmail && password != LoggedPass {
				fmt.Println("logeado como usuario")
				return true
			}

			return false
		}

		login("jfjslda", 84023)
		return
	}

	fmt.Println("no logged")

}
*/
