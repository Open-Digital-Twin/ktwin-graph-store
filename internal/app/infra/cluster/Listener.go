package cluster

import (
	"fmt"

	"time"

	twingraph "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/pkg/log"
)

func NewClusterListener(cacheConnection cache.CacheConnection, logger log.Logger) ClusterListener {
	return &clusterListener{
		logger:          logger,
		cacheConnection: cacheConnection,
		clusterClient:   NewClusterClient(logger),
	}
}

type ClusterListener interface {
	Listen(waitTime int)
}

type clusterListener struct {
	logger          log.Logger
	clusterClient   ClusterClient
	cacheConnection cache.CacheConnection
}

func (c *clusterListener) Listen(waitTime int) {
	c.logger.Info(fmt.Sprintf("Listen wait time: %d\n", waitTime))
	go func() {
		for {
			c.logger.Info("Listening for twin instances\n")
			c.listenTwinInstances()
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}()
}

func (c *clusterListener) listenTwinInstances() {
	result, err := c.clusterClient.GetResources("/apis/dtd.ktwin/v0", "ktwin", "twininstances")
	if err == nil {
		container := twingraph.InitializeTwinGraphContainer(c.cacheConnection)
		container.Controller.UpdateTwinGraph(result)
	} else {
		c.logger.Error(fmt.Sprintf("Error getting resources from cluster: %v", err))
	}
}
