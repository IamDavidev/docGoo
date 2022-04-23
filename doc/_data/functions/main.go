package main

import "fmt"

var EMAIL_LOGGED_IN = "dav@gmail.com"
var PASSWORD_LOGGED_IN = "dav123"

func main() {

	fmt.Println("Welcome to the login page")

	// login(EMAIL_LOGGED_IN, PASSWORD_LOGGED_IN)

	// apro := promedio(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

}

// for range loop || return the promeio

func promedio(calificaciones ...float64) (promedio float64) {

	suma := 0.0
	cal := 0

	for index, item := range calificaciones {

		cal = index
		suma = suma + item

	}

	fmt.Println(cal)
	promedio = suma

	return

}

func login(email string, password string) bool {

	var isEmail bool = email == EMAIL_LOGGED_IN
	var isPassword bool = password == PASSWORD_LOGGED_IN

	if !isEmail {
		fmt.Println("Email is not correct")
		return false
	}

	if !isPassword {
		fmt.Println("Password is not correct")
		return false
	}

	fmt.Println("credentials  correct")

	return false
}
