package magazine

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address // A struct field that is itself a struct type
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Author struct {
	Book string

	// Anonymous struct fields: struct fields that have no name of their own, just a type
	Address // The field's type name is used as if it were the name of the field
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
