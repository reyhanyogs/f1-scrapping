package main

import (
	"f1-scrapping/handler"
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("div#driver-standings > table.data-table > tbody", handler.DriverConstructor)
	c.OnHTML("div#constructor-standings > table.data-table > tbody", handler.TeamConstructor)
	c.OnHTML("div.section > table.data-table > tbody", handler.RaceWinner)
	c.OnHTML("div.lead_laps > table.data-table > tbody", handler.DriverTotalLeadLaps)

	var year string
	fmt.Print("Enter The Desired F1 Season to Scrape: ")
	fmt.Scanln(&year)
	fmt.Printf("F1 %s Season Scrape Result\n\n", year)

	c.Visit(scrapeUrl(year))

	fmt.Printf("%s Driver Constructors Standings\n", year)
	fmt.Printf("{Position, Driver, Team, Points}\n")
	for _, data := range handler.DriversConstructor {
		fmt.Printf("{%s, %s, %s, %s}\n", data.Position, data.Name, data.Team, data.Point)
	}

	fmt.Printf("\n%s Team Constructors Standings\n", year)
	fmt.Printf("{Position, Team, Points}\n")
	for _, data := range handler.TeamsConstructor {
		fmt.Printf("{%s, %s, %s}\n", data.Position, data.TeamName, data.Point)
	}

	fmt.Printf("\n%s Track Winner\n", year)
	fmt.Printf("{Date, Track, Winner}\n")
	for _, data := range handler.RacesWinner {
		fmt.Printf("{%s, %s, %s}\n", data.Date, data.Track, data.Winner)
	}

	fmt.Printf("\nDriver Total Lead Laps\n")
	fmt.Printf("{Position, Driver, Total Laps}\n")
	for _, data := range handler.DriversLeadLaps {
		fmt.Printf("{%s, %s, %s}\n", data.Position, data.DriverName, data.TotalLaps)
	}
}

func scrapeUrl(year string) string {
	return "https://pitwall.app/seasons/" + year + "-formula-1-world-championship"
}
