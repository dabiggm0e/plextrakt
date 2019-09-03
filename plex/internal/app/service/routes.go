package service

import "github.com/dabiggm0e/plextrakt/common/router"

// Initialize our routes
var routes = router.Routes{
	router.Route{
		"ParseEvent",
		"POST",
		"/plex/events",
		ParsePlexEvent,
		true,
	},
}
