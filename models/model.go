package models

type JobSummary struct {
	Fault  string `json:fault`
	Status string `json:status`
	Car    string `json:car`
	Id     int    `json:id`
}

type Person struct {
	Contact string `json:contact`
	Name    string `json:name`
	Address string `json:address`
}

type Car struct {
	CarNum   string `json:carNum`
	CarName  string `json:carName`
	Color    string `json:color`
	CarOwner string `json:carOwner`
}

type JobSummaryData struct {
	Fault  string
	Status string
	Id     int
}

type CarData struct {
	CarNum  string
	CarName string
	Color   string
	Jobs    []JobSummaryData
}

type PersonData struct {
	Contact string
	Name    string
	Address string
	Cars    []CarData
}
