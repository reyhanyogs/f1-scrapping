package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type driver struct {
	position string
	name     string
	car      string
	point    string
}

type constructor struct {
	position string
	name     string
	point    string
}

type raceWinner struct {
	date   string
	track  string
	winner string
}

type driverLeadLaps struct {
	position  string
	driver    string
	totalLaps string
}

func main() {
	c := colly.NewCollector()
	drivers := []driver{}
	constructors := []constructor{}
	raceWinners := []raceWinner{}
	driversLeadLaps := []driverLeadLaps{}

	c.OnHTML("div#driver-standings > table.data-table > tbody", func(row *colly.HTMLElement) {
		row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
			driver := driver{}
			driver.position = cell.ChildText("td:nth-child(1)")
			driver.name = cell.ChildText("td:nth-child(2) a")
			driver.car = cell.ChildText("td:nth-child(3)")
			driver.point = cell.ChildText("td:nth-child(6)")
			drivers = append(drivers, driver)
		})
		fmt.Println(drivers)
		fmt.Println()
	})

	c.OnHTML("div#constructor-standings > table.data-table > tbody", func(row *colly.HTMLElement) {
		row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
			constructor := constructor{}
			constructor.position = cell.ChildText("td:nth-child(1)")
			constructor.name = cell.ChildText("td:nth-child(2) a")
			constructor.point = cell.ChildText("td:nth-child(5)")
			constructors = append(constructors, constructor)
		})
		fmt.Println(constructors)
		fmt.Println()
	})

	c.OnHTML("div.section > table.data-table > tbody", func(row *colly.HTMLElement) {
		row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
			raceWinner := raceWinner{}
			raceWinner.date = cell.ChildText("td:nth-child(1)")
			raceWinner.track = cell.ChildText("td:nth-child(2) a")
			splitDriverName := strings.Split(cell.ChildText("td:nth-child(3)"), " ")
			raceWinner.winner = strings.Join(splitDriverName[1:], " ")
			raceWinners = append(raceWinners, raceWinner)
		})
		fmt.Println(raceWinners)
		fmt.Println()
	})

	c.OnHTML("div.lead_laps > table.data-table > tbody", func(row *colly.HTMLElement) {
		row.ForEach("tr", func(_ int, cell *colly.HTMLElement) {
			driverLeadLaps := driverLeadLaps{}
			driverLeadLaps.position = cell.ChildText("td:nth-child(1)")
			driverLeadLaps.driver = cell.ChildText("td:nth-child(2) a")
			driverLeadLaps.totalLaps = cell.ChildText("td:nth-child(3)")
			driversLeadLaps = append(driversLeadLaps, driverLeadLaps)
		})
		fmt.Println(driversLeadLaps)
		fmt.Println()
	})

	c.Visit(scrapeUrl("2020"))
}

func scrapeUrl(year string) string {
	return "https://pitwall.app/seasons/" + year + "-formula-1-world-championship"
}
