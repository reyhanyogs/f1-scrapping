package handler

import (
	"f1-scrapping/model"
	"strings"

	"github.com/gocolly/colly/v2"
)

var (
	DriversConstructor []model.Driver
	TeamsConstructor   []model.Team
	RacesWinner        []model.RaceWinner
	DriversLeadLaps    []model.DriverLeadLaps
)

func DriverConstructor(row *colly.HTMLElement) {
	row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
		driver := model.Driver{}
		driver.Position = cell.ChildText("td:nth-child(1)")
		driver.Name = cell.ChildText("td:nth-child(2) a")
		driver.Team = cell.ChildText("td:nth-child(3)")
		driver.Point = cell.ChildText("td:nth-child(6)")
		DriversConstructor = append(DriversConstructor, driver)
	})
}

func TeamConstructor(row *colly.HTMLElement) {
	row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
		constructor := model.Team{}
		constructor.Position = cell.ChildText("td:nth-child(1)")
		constructor.TeamName = cell.ChildText("td:nth-child(2) a")
		constructor.Point = cell.ChildText("td:nth-child(5)")
		TeamsConstructor = append(TeamsConstructor, constructor)
	})
}

func RaceWinner(row *colly.HTMLElement) {
	row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
		raceWinner := model.RaceWinner{}
		raceWinner.Date = cell.ChildText("td:nth-child(1)")
		raceWinner.Track = cell.ChildText("td:nth-child(2) a")
		splitDriverName := strings.Split(cell.ChildText("td:nth-child(3)"), " ")
		raceWinner.Winner = strings.Join(splitDriverName[1:], " ")
		RacesWinner = append(RacesWinner, raceWinner)
	})
}

func DriverTotalLeadLaps(row *colly.HTMLElement) {
	row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
		driverLeadLaps := model.DriverLeadLaps{}
		driverLeadLaps.Position = cell.ChildText("td:nth-child(1)")
		driverLeadLaps.DriverName = cell.ChildText("td:nth-child(2) a")
		driverLeadLaps.TotalLaps = cell.ChildText("td:nth-child(3)")
		DriversLeadLaps = append(DriversLeadLaps, driverLeadLaps)
	})
}
