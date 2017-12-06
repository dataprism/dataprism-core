package main

import (
	"github.com/dataprism/dataprism-commons/core"
	"github.com/dataprism/dataprism-commons/api"
)

func CreateRoutes(platform *core.Platform, API *api.Rest) {
	evaluationManager := core.NewEvaluationManager(platform)
	evaluationRouter := core.NewEvaluationRouter(evaluationManager)
	API.RegisterGet("/v1/evaluations/{id}", evaluationRouter.Get)
	API.RegisterGet("/v1/evaluations/{id}/events", evaluationRouter.Events)

	nodeManager := core.NewNodeManager(platform)
	nodeRouter := core.NewNodeRouter(nodeManager)
	API.RegisterGet("/v1/nodes", nodeRouter.List)
	API.RegisterGet("/v1/nodes/{id}", nodeRouter.Get)
}
