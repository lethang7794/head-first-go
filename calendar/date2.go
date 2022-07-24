package calendar

import "errors"

type Date2 struct {
	year  int
	month int
	day   int
}

func (d *Date2) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date2) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date2) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

// Use a pointer receiver type for consistency with the setter methods
func (d *Date2) Year() int { // The method name is same name as the field (but capitalized so it's exported)
	return d.year
}

func (d *Date2) Month() int {
	return d.month
}

func (d *Date2) Day() int {
	return d.day
}
