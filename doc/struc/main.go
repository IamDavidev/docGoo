package main

import (
	"fmt"
	"time"

	note "struc/note"
)

type NoteCreator struct {
	note.Note
}

func main() {

	// primer forma de crear un nota

	firstNote := new(NoteCreator)
	firstNote.NewNote(1, "First Note", false, time.Now())
	fmt.Println("first note =>",firstNote.Note)

	//
	var secondNote NoteCreator
	secondNote.NewNote(2, "Second Note", true, time.Now())
	fmt.Println("second note =>",secondNote.Note)
}
