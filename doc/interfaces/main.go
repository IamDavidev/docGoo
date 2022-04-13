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
}

// var globales empiezan con mayuscula

// var locales empiezan con minuscula

func main() {

	var firstArticle Article = Article{
		Title:  "First Article",
		Author: "Daniel",
		Date:   time.Now(),
	}

	var secondArticle Article = Article{
		Title:  "Second Article",
		Author: "Daniel",
		Date:   time.Now(),
	}
	printDescription(secondArticle)
	printDescription(firstArticle)
}

func printDescription(method Description) {

	fmt.Println(method.RenderDescription())
}

func (this Article) RenderDescription() string {
	return fmt.Sprintf("The %q article(%v) was written by %s .", this.Title, this.Date, this.Author)
}
