package datafile

//Subscriber information
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	// HomeAddress Address

	//Anonymous field
	Address
}

//Employee information
type Employee struct {
	Name   string
	Salary float64
	// HomeAddress Address

	//Anonymous field
	Address
}

//Address information
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
