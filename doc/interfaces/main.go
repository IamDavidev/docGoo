package main

import (
	"fmt"
	"time"
)

type DemoInterface interface {

	// se define el comportamiento de la interfaz
	// ` Que puede hacer cada este type `

	String() string
}

type Article struct {
	Title  string
	Author string
	Date   time.Time
}

type Description interface {
	RenderDescription() string
	RenderAuthor() string
}

type Author interface {
	RenderAuthor() string
}

// var globales empiezan con mayuscula

// var locales empiezan con minuscula

func main() {

	var firstArticle Article = Article{

		Title:  "First Article",
		Author: "Daniel",
		Date:   time.Now(), // retorna la fecha actual

	}

	var secondArticle Article = Article{

		Title:  "Second Article",
		Author: "Daniel",
		Date:   time.Now(),
	}
	printDescription(secondArticle)
	prinAuthor(firstArticle)
}

func printDescription(propArticle Description) {

	fmt.Println(propArticle.RenderDescription())
	fmt.Println(propArticle.RenderAuthor())

}

func prinAuthor(propArticle Author) {

	fmt.Println(propArticle.RenderAuthor())

}

func (this Article) RenderDescription() string {
	return fmt.Sprintf("The %q article(%v) was written by %s .", this.Title, this.Date, this.Author)
}

func (this Article) RenderAuthor() string {
	return fmt.Sprintf(" written by %s .", this.Author)
}
