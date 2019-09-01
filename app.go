package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dabiggm0e/plextrakt/plex"
)

func main() {
	file, err := ioutil.ReadFile("plex/samples/movie-scrobble.json")
	if err != nil {
		fmt.Print(err)
	}

	//	str := string(file)
	//	fmt.Print(str)

	webhook := plex.Webhook{}
	err = json.Unmarshal([]byte(file), &webhook)

	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
	}
	fmt.Printf("%+v", webhook)
}
