package main

import (
	"f1-scrapping/handler"
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
)

func main() {
	var (
		year       string
		counter    = 0
		c          = colly.NewCollector()
		xlFile     = excelize.NewFile()
		sheetNames = []string{"Driver Constructor Standing", "Team Constructor Standings", "Track Winner", "Driver Total Lead Laps"}
	)

	defer func() {
		if err := xlFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	c.OnHTML("div#driver-standings > table.data-table > tbody", handler.DriverConstructor)
	c.OnHTML("div#constructor-standings > table.data-table > tbody", handler.TeamConstructor)
	c.OnHTML("div.section > table.data-table > tbody", handler.RaceWinner)
	c.OnHTML("div.lead_laps > table.data-table > tbody", handler.DriverTotalLeadLaps)

	fmt.Print("Enter The Desired F1 Season to Scrape: ")
	fmt.Scanln(&year)
	fmt.Printf("F1 %s Season Scrape Result\n\n", year)

	c.Visit(scrapeUrl(year))

	// Creating Sheet
	for counter <= 3 {
		_, err := xlFile.NewSheet(sheetNames[counter])
		if err != nil {
			panic(err)
		}
		counter++
	}
	xlFile.DeleteSheet("Sheet1")

	// Driver Constructors Standings Data Setup
	fmt.Printf("%s Driver Constructors Standings\n", year)
	fmt.Printf("{Position, Driver, Team, Points}\n")
	xlFile.SetCellValue(sheetNames[0], "A1", "Position")
	xlFile.SetCellValue(sheetNames[0], "B1", "Driver")
	xlFile.SetCellValue(sheetNames[0], "C1", "Team")
	xlFile.SetCellValue(sheetNames[0], "D1", "Points")
	for i, data := range handler.DriversConstructor {
		xlFile.SetCellValue(sheetNames[0], fmt.Sprintf("A%d", i+2), data.Position)
		xlFile.SetCellValue(sheetNames[0], fmt.Sprintf("B%d", i+2), data.Name)
		xlFile.SetCellValue(sheetNames[0], fmt.Sprintf("C%d", i+2), data.Team)
		xlFile.SetCellValue(sheetNames[0], fmt.Sprintf("D%d", i+2), data.Point)
		fmt.Printf("{%s, %s, %s, %s}\n", data.Position, data.Name, data.Team, data.Point)
	}

	// Team Constructor Standings Data Setup
	fmt.Printf("\n%s Team Constructor Standings\n", year)
	fmt.Printf("{Position, Team, Points}\n")
	xlFile.SetCellValue(sheetNames[1], "A1", "Position")
	xlFile.SetCellValue(sheetNames[1], "B1", "Team Name")
	xlFile.SetCellValue(sheetNames[1], "C1", "Points")
	for i, data := range handler.TeamsConstructor {
		xlFile.SetCellValue(sheetNames[1], fmt.Sprintf("A%d", i+2), data.Position)
		xlFile.SetCellValue(sheetNames[1], fmt.Sprintf("B%d", i+2), data.TeamName)
		xlFile.SetCellValue(sheetNames[1], fmt.Sprintf("C%d", i+2), data.Point)
		fmt.Printf("{%s, %s, %s}\n", data.Position, data.TeamName, data.Point)
	}

	// Track Winner Data Setup
	fmt.Printf("\n%s Track Winner\n", year)
	fmt.Printf("{Date, Track, Winner}\n")
	xlFile.SetCellValue(sheetNames[2], "A1", "Date")
	xlFile.SetCellValue(sheetNames[2], "B1", "Track")
	xlFile.SetCellValue(sheetNames[2], "C1", "Winner")
	for i, data := range handler.RacesWinner {
		xlFile.SetCellValue(sheetNames[2], fmt.Sprintf("A%d", i+2), data.Date)
		xlFile.SetCellValue(sheetNames[2], fmt.Sprintf("B%d", i+2), data.Track)
		xlFile.SetCellValue(sheetNames[2], fmt.Sprintf("C%d", i+2), data.Winner)
		fmt.Printf("{%s, %s, %s}\n", data.Date, data.Track, data.Winner)
	}

	// Driver Total Lead Laps Data Setup
	fmt.Printf("\nDriver Total Lead Laps\n")
	fmt.Printf("{Position, Driver, Total Laps}\n")
	xlFile.SetCellValue(sheetNames[3], "A1", "Position")
	xlFile.SetCellValue(sheetNames[3], "B1", "Driver")
	xlFile.SetCellValue(sheetNames[3], "C1", "Total Laps")
	for i, data := range handler.DriversLeadLaps {
		xlFile.SetCellValue(sheetNames[3], fmt.Sprintf("A%d", i+2), data.Position)
		xlFile.SetCellValue(sheetNames[3], fmt.Sprintf("B%d", i+2), data.DriverName)
		xlFile.SetCellValue(sheetNames[3], fmt.Sprintf("C%d", i+2), data.TotalLaps)
		fmt.Printf("{%s, %s, %s}\n", data.Position, data.DriverName, data.TotalLaps)
	}

	// Save Excel File
	if err := xlFile.SaveAs("scrape-result.xlsx"); err != nil {
		panic(err)
	}
}

func scrapeUrl(year string) string {
	return "https://pitwall.app/seasons/" + year + "-formula-1-world-championship"
}
