package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("leer archivos en go ")

	document, err := os.Open("./index.txt")
	if err != nil {
		fmt.Println("error al abrir el archivo")
	}
	scanner := bufio.NewScanner(document)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("newLine = ", line)
		document.Close()
	}
	// crear un archivo

	newFile, err := os.Create("./index.js")
	if err != nil {
		fmt.Println("error al abrir el archivo")
	}

	newFile.WriteString("document.write('hola mundo')")

}
