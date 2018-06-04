package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/tealeg/xlsx"
)

type Artist struct {
	Name string
}

func scrapeWebsite(website string) (artists []Artist) {

	c := colly.NewCollector()

	c.OnHTML("span", func(e *colly.HTMLElement) {
		artists = append(artists, Artist{
			Name: e.Text,
		})
	})

	c.Visit(website)
	return artists
}

func writeToExcel(artists []Artist, fileName string) error {

	f := xlsx.NewFile()
	sheet, err := f.AddSheet("Artists")
	if err != nil {
		return err
	}

	for _, a := range artists {
		row := sheet.AddRow()

		cell := row.AddCell()
		cell.Value = a.Name
	}

	if err := f.Save(fileName); err != nil {
		return err
	} else {
		fmt.Println("successfully generated:", fileName)
		return nil
	}
}

func main() {

	website := "https://www.boomfestival.org/boom2018/program/alchemy-circle/"
	fileName := "alchemy_circle.xlsx"
	artists := scrapeWebsite(website)
	writeToExcel(artists, fileName)

	website = "https://www.boomfestival.org/boom2018/program/chill-out-gardens/music/"
	fileName = "chill_out_gardens.xlsx"
	artists = scrapeWebsite(website)
	writeToExcel(artists, fileName)
}
