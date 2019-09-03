package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/alexflint/go-arg"
	"github.com/dabiggm0e/plextrakt/plex/cmd"
)

var appName = "plexService"

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Starting %v\n", appName)

	// Initialize config struct and populate it froms env vars and flags.
	cfg := cmd.DefaultConfiguration()
	arg.MustParse(cfg)

	//mc := initializeMessaging(cfg)

	client := &http.Client{}
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}

	client.Transport = transport
	_ = client

	h := service.NewHandler(mc, client)

	s := service.NewServer(cfg, h)
	s.SetupRoutes()

	handleSigterm(func() {
		mc.Close()
		s.Close()
	})

	s.Start()

	// load a sample webhook json file
	/*file, err := ioutil.ReadFile("plex/internal/app/testing/samples/show-scrobble2.json")
		if err != nil {
			fmt.Print(err)
		}

		//	str := string(file)
		//	fmt.Print(str)

		// make sure it has been unmarshaled successfully
		event := model.Event{}
		err = json.Unmarshal([]byte(file), &event)

		if err != nil {
			log.Printf("Error unmarshaling: %v", err)
		}
		fmt.Printf("%+v", event)
	}
	*/

	/*func initializeMessaging(cfg *cmd.Config) *messaging.AmqpClient {

	 */
}

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
