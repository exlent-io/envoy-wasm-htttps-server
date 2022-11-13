package main

import (
	"context"
	"flag"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/processor"
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/server"
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/watcher"
	log "github.com/sirupsen/logrus"
)

var (
	l log.FieldLogger

	configFileName       string
	watchVersionFileName string
	port                 uint
	basePort             uint
	mode                 string

	nodeID string
)

func init() {
	l = log.New()
	log.SetLevel(log.DebugLevel)

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	// Define the directory to store management server configuration
	flag.StringVar(&configFileName, "configFileName", "config/config.yaml", "Config File Path")

	// Define the directory to watch for VERSION file
	flag.StringVar(&watchVersionFileName, "watchVersionFileName", "version.txt", "The file used to notify the server that the config should be reload")
}

func main() {
	flag.Parse()

	// Create a cache
	cache := cache.NewSnapshotCache(true, cache.IDHash{}, l)

	// Create a processor
	proc := processor.NewProcessor(
		cache, nodeID, log.WithField("context", "processor"), configFileName)

	// Create initial snapshot from file
	proc.ProcessFile(watcher.NotifyMessage{
		Operation: watcher.Create,
		FilePath:  watchVersionFileName,
	})

	// Notify channel for file system events
	notifyCh := make(chan watcher.NotifyMessage)

	go func() {
		// Watch for file changes
		// watcher.Watch(watchDirectoryFileName, notifyCh)
		watcher.Watch(watchVersionFileName, notifyCh)
	}()

	go func() {
		// Run the xDS server
		ctx := context.Background()
		srv := serverv3.NewServer(ctx, cache, nil)
		server.RunServer(ctx, srv, port)
	}()

	for {
		select {
		case msg := <-notifyCh:
			log.Printf("receive msg: %+v", msg)
			proc.ProcessFile(msg)
		}
	}
}
