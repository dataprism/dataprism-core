package main

import (
	"github.com/dataprism/dataprism-commons/api"
	"github.com/sirupsen/logrus"
	"flag"
	"os"
	"github.com/dataprism/dataprism-commons/core"
	"github.com/dataprism/dataprism-commons/plugins"
	"github.com/dataprism/dataprism-sync/sync"
	core2 "github.com/dataprism/dataprism-core/core"
)

func main() {
	var bind = flag.String("b", "0.0.0.0:6100", "the bind address, including the port on which the api will listen")

	flag.Parse()

	// -- load the plugins
	registry := plugins.NewDataprismPluginRegistry()
	registry.Add(&core2.DataprismPlugin{})
	registry.Add(&sync.DataprismPlugin{})

	// -- initialize the platform
	platform, err := core.InitializePlatform()
	if err != nil { logrus.Fatal(err) }

	// -- create the api
	API := api.CreateAPI(*bind)

	// -- route the api endpoints
	for _, p := range registry.Plugins() {
		p.CreateRoutes(platform, API)
	}

	// -- start serving the api
	err = API.Start()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Info("API listening on http://" + *bind + "/v1")
}