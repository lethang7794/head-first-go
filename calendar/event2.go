package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event2 struct {
	title string // changed to unexported field
	Date2
}

func (e Event2) Title() string {
	return e.title
}

func (e *Event2) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
