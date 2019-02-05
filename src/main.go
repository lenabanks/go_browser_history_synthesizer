//-- Package Declaration -----------------------------------------------------------------------------------------------
package main

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"github.com/JustonDavies/go_activity_synthesizer/src/browsers"
	"github.com/JustonDavies/go_activity_synthesizer/src/conf"
	"github.com/cheggaaa/pb"
	"log"
	"math/rand"
	"time"
)

//-- Constants ---------------------------------------------------------------------------------------------------------

//-- Structs -----------------------------------------------------------------------------------------------------------

//-- Exported Functions ------------------------------------------------------------------------------------------------
func main() {
	//-- Log nice output ----------
	var start = time.Now().Unix()
	log.Println(`Starting task...`)

	//-- Perform task ----------
	var thingme = browsers.Open()
	rand.Seed(time.Now().UnixNano())

	if len(thingme) < 1 {
		panic(`unable to open any supported browsers, aborting...`)
	}

	log.Println(`Injecting history...`)
	var historyProgress = pb.StartNew(len(conf.ActivityItems))
	for _, item := range conf.ActivityItems {
		var browser = thingme[rand.Intn(len(thingme))]
		var item = browsers.History{
			Name:        item.Name,
			URL:         item.URL,
			Visits:      rand.Intn(conf.MaximumVisits),
			VisitWindow: conf.DefaultDuration,
		}

		if err := browser.InjectHistory(item); err != nil {
			log.Printf("unable to inject history item for: \n\tURL: '%s' \n\tError: '%s'", item.URL, err)
		}
		historyProgress.Increment()
	}

	log.Println(`Injecting bookmarks...`)
	var bookmarkProgress = pb.StartNew(len(conf.ActivityItems))
	for _, item := range conf.ActivityItems {
		var browser = thingme[rand.Intn(len(thingme))]
		var item = browsers.Bookmark{
			Name:         item.Name,
			URL:          item.URL,
			CreateWindow: conf.DefaultDuration,
		}

		if err := browser.InjectBookmark(item); err != nil {
			log.Printf("unable to inject bookmark item for: \n\tURL: '%s' \n\tError: '%s'", item.URL, err)
		}
		bookmarkProgress.Increment()
	}

	browsers.Close(thingme)
	//-- Log nice output ----------
	log.Printf(`Task complete! It took %d seconds`, time.Now().Unix()-start)
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
