package main

import (
	"log"
	"fmt"
	"strconv"
)

type Fight struct {
	Fighter1Name string
	Fighter2Name string
	Rounds  {}
	Fighter1Age int
	Fighter2Age  int
	Fighter1Height int
	Fighter2Height int
	Fighter1Reach string
	Fighter2Reach string
	Fighter1Stance string
	Fighter2Stance string
	Fighter1Form {}
	Fighter2Form {}
	Fighter1Wins int
	Fighter2Wins int
	Fighter1Byko int 
	Fighter2Byko int
	Fighter1K0% int 
	Fighter2K0% int 
    Fighter1Loses int
    Fighter2Loses int
    Fighter1Draws int
	Fighter2Draws int
	Fighter1Debut int
	Fighter2Debut int
	Fighter1The-Ring int
	Fighter2The-Ring int
	Fighter1WBA int
	Fighter2WBA int
	Fighter1WBC int
	Fighter2WBC int
	Fighter1WBO int
	Fighter2WBO int
	Fighter1IBF int
	Fighter2IBF int

}

type Fighters struct {
}

// type SubFights Fight {

// }

type FightResults struct {
	Date {}
	Participant1Name string
	Participant2Name string
	FightInfo {}
	Event-Place  Event_Place 
	UnderCard_Events {}
}

type EventPlace struct {
	Stadium_Name string
	Country_Name string
}




type UpcomingFights struct {
}

func main() {
	// slices
	allFights := make([]Fight, 0)
	FightersName := make([]Fighters, 0)
	FightResults := make([]FightResults, 0)
	EventPlace := make([]EventPlace, 0)


	// create a new collector
	collector := colly.NewCollector(
        colly.AllowedDomains("https://box.live/fight-results/", "https://box.live/upcoming-fights-schedule/"),
    )


	// from inspect 
	collector.OnHTML("", func(element *colly.HTMLElement) {

       boxId, err := strconv.Atoi(element.Attr("id"))
	   // error handling
        if err != nil {
            log.Println("Could not get id") // id is string and convert to int
        }

        fightDesc := element.Text

      fight := Fights{
            ID:          fightId,
            Description: fightDesc,
        }

        allFights = append(allFights, fight)
    })

	
	//  Begin crawling and scraping
collector.OnRequest(func(request *colly.Request) {
	fmt.Println("Visiting", request.URL.String())
})

collector.Visit("https://box.live/fight-results/", "https://box.live/upcoming-fights-schedule/")

// save in json format
writeJSON(allFacts)


}

// Saving our data to JSON

func writeJSON(data []Fights) {
    file, err := json.MarshalIndent(data, "", " ")
    if err != nil {
        log.Println("Unable to create json file")
        return
    }

    _ = ioutil.WriteFile("fights.json", file, 0644)
}