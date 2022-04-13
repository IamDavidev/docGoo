package main

import (
	"fmt"
)

type Estado struct {
	Nombre string
	Idioma string
}

type Pais struct {
	Nombre    string
	Poblacion uint64
	Capital   string
	// EstadosPrincipales []Estado
}

var Paises = make([]Pais, 3, 5)

func main() {
	fmt.Println("slices and structs")

	var creted = "si"
	var name, capital string
	var poblacion uint64

	for {
		fmt.Println("1. Crear Pais")

		fmt.Println("Cual es el nombre del pais que quieres crear?")
		fmt.Scanln(&name)

		fmt.Println("Cual es la capital del pais que quieres crear?")
		fmt.Scanln(&capital)

		fmt.Println("Cual es la poblacion aproximado del pais que quieres crear?")
		fmt.Scanln(&poblacion)

		NuevoPais(name, poblacion, capital)

		fmt.Println("Desea crear otro pais? (si/no)")
		fmt.Scanln(&creted)

		if creted != "si" {
			break
		}

	}
}

// name string , poblacion uint64, capital string , estadosPrincipales []Estado

func NuevoPais(nombre string, poblacion uint64, capital string) {

	var crearPais Pais = Pais{
		Nombre:    nombre,
		Poblacion: poblacion,
		Capital:   capital,
		// EstadosPrincipales: estadosPrincipales,
	}

	Paises = append(Paises, crearPais)
	fmt.Println("Pasis Creado ")
	fmt.Println(Paises)

}
