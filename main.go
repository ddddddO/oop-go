package main

import (
	"github.com/ddddddO/oop-go/curator"
	"github.com/ddddddO/oop-go/curator/api"
	"github.com/ddddddO/oop-go/curator/scrape"
)

func main() {
	curators := []curator.Curator{
		api.NewTwitterCurator("Twitter", "twid", "twipass"),
		api.NewGurunaviCurator("Gurunavi", "guruid", "gurupass"),
		scrape.NewTabelogCurator("Tabelog"),
		api.NewSpotifyCurator("Spotify", "spoid", "spopass"),
	}

	if err := curator.Run(curators); err != nil {
		panic(err)
	}
}
