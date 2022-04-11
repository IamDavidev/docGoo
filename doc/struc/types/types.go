package types

import "time"

type Note struct {
	Id        int
	Title     string
	Completed bool
	Date      time.Time
}

func (self *Note) NewNote(id int, title string, completed bool, date time.Time) {
	self.Id = id
	self.Title = title
	self.Completed = completed
	self.Date = date

}
