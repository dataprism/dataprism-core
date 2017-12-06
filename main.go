package main

import (
	"github.com/dataprism/dataprism-commons/api"
	"github.com/sirupsen/logrus"
	"flag"
	"strconv"
	"os"
	"github.com/dataprism/dataprism-commons/core"
)

func main() {
	var port = flag.Int("p", 6000, "the port of the dataprism core rest api")

	flag.Parse()

	platform, err := core.InitializePlatform()
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(1)
	}

	// -- create the api
	API := api.CreateAPI("0.0.0.0:" + strconv.Itoa(*port))

	// -- route the api endpoints
	CreateRoutes(platform, API)

	// -- start serving the api
	err = API.Start()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Info("API listening on http://0.0.0.0:" + strconv.Itoa(*port) + "/v1")
}