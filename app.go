package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dabiggm0e/plextrakt/plex"
)

func main() {
	// load a sample webhook json file
	file, err := ioutil.ReadFile("plex/samples/show-scrobble2.json")
	if err != nil {
		fmt.Print(err)
	}

	//	str := string(file)
	//	fmt.Print(str)

	// make sure it has been unmarshaled successfully
	webhook := plex.Webhook{}
	err = json.Unmarshal([]byte(file), &webhook)

	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
	}
	fmt.Printf("%+v", webhook)

}
