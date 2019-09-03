package service

import "github.com/dabiggm0e/plextrakt/common/router"

// Initialize our routes
var routes = router.Routes{
	router.Route{
		Name:        "ParseEvent",
		Method:      "POST",
		Pattern:     "/plex/events",
		HandlerFunc: ParsePlexEvent,
		Monitor:     true,
	},
}
