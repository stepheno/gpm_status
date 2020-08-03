// Google Play Music Status Bar parser
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	color = flag.String("color", "green", "Status output color")
)

type Song struct {
	Title    string
	Artist   string
	Album    string
	AlbumArt string
}

type Rating struct {
	Liked    bool
	Disliked bool
}

type Progress struct {
	Current int
	Total   int
}

type Status struct {
	Playing    bool
	Song       Song
	SonyLyrics string
	Shuffle    string
	Repeat     string
	Volume     int
	Rating     Rating
	Progress   Progress
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println(os.Stderr, "Usage: gpm_status FILE")
		os.Exit(1)
	}

	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(os.Stderr, "Error reading status file: %v\n", err)
		os.Exit(1)
	}

	bytes, _ := ioutil.ReadAll(file)
	var result Status

	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(os.Stderr, "Error parsing json: %v", err)
		os.Exit(1)
	}

	song := result.Song
	if result.Playing {
		fmt.Fprintf(os.Stdout, "%s - %s", song.Artist, song.Title)
	} else {
		fmt.Printf("Paused")
	}
}
