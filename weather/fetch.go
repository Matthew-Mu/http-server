package weather

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
)

type Weather struct {
	Date,
	Day,
	TempLow,
	TempHigh,
	Rain,
	Evap,
	Sun,
	WindDir,
	MaxWindSpd,
	Time,
	AMTemp,
	AMRH,
	AMCloud,
	AMWindDir,
	AMMaxWindSpd,
	AMPressure,
	PMTemp,
	PMRH,
	PMCloud,
	PMWindDir,
	PMMaxWindSpd,
	PMPressure string
}

func Fetch() []byte {

	var weatherData []Weather

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
		//fmt.Println("Response: \n", r.Body)
		//fmt.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Find and visit all links
	c.OnHTML("tbody:nth-child(odd)", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			weatherTable := Weather{
				Date:         el.ChildText("th"),
				Day:          el.ChildText("td:nth-child(2)"),
				TempLow:      el.ChildText("td:nth-child(3)"),
				TempHigh:     el.ChildText("td:nth-child(4)"),
				Rain:         el.ChildText("td:nth-child(5)"),
				Evap:         el.ChildText("td:nth-child(6)"),
				Sun:          el.ChildText("td:nth-child(7)"),
				WindDir:      el.ChildText("td:nth-child(8)"),
				MaxWindSpd:   el.ChildText("td:nth-child(9)"),
				Time:         el.ChildText("td:nth-child(10)"),
				AMTemp:       el.ChildText("td:nth-child(11)"),
				AMRH:         el.ChildText("td:nth-child(12)"),
				AMCloud:      el.ChildText("td:nth-child(13)"),
				AMWindDir:    el.ChildText("td:nth-child(14)"),
				AMMaxWindSpd: el.ChildText("td:nth-child(15)"),
				AMPressure:   el.ChildText("td:nth-child(16)"),
				PMTemp:       el.ChildText("td:nth-child(17)"),
				PMRH:         el.ChildText("td:nth-child(18)"),
				PMCloud:      el.ChildText("td:nth-child(19)"),
				PMWindDir:    el.ChildText("td:nth-child(20)"),
				PMMaxWindSpd: el.ChildText("td:nth-child(21)"),
				PMPressure:   el.ChildText("td:nth-child(22)"),
			}
			weatherData = append(weatherData, weatherTable)
		})
	})

	c.Request("GET",
		"http://www.bom.gov.au/climate/dwo/IDCJDW2124.latest.shtml",
		nil,
		nil,
		http.Header{"User-Agent": []string{"fhdoxsqxlsbdbebdn"}},
	)

	content, err := json.Marshal(weatherData)
	if err != nil {
		fmt.Println(err.Error())
	}

	return content
}
