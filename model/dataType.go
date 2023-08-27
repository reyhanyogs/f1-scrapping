package model

type Driver struct {
	Position string
	Name     string
	Team     string
	Point    string
}

type Team struct {
	Position string
	TeamName string
	Point    string
}

type RaceWinner struct {
	Date   string
	Track  string
	Winner string
}

type DriverLeadLaps struct {
	Position   string
	DriverName string
	TotalLaps  string
}
