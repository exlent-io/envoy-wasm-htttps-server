package processor

import (
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/resources"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/adscache"
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/watcher"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	cache  cache.SnapshotCache
	nodeID string

	// snapshotVersion holds the current version of the snapshot.
	snapshotVersion int64

	logrus.FieldLogger

	// xdsCache xdscache.XDSCache
	adsCache adscache.ADSCache
}

func NewProcessor(cache cache.SnapshotCache, nodeID string, log logrus.FieldLogger) *Processor {
	return &Processor{
		cache:           cache,
		nodeID:          nodeID,
		snapshotVersion: rand.Int63n(1000),
		FieldLogger:     log,
		//xdsCache: xdscache.XDSCache{
		//	Listeners:    make(map[string]resources.Listener),
		//	Clusters:     make(map[string]resources.Cluster),
		//	Routes:       make(map[string]resources.Route),
		//	Endpoints:    make(map[string]resources.Endpoint),
		//	TLSListeners: make(map[string]resources.TLSListener),
		//	Secrets:      make(map[string]resources.Secret),
		//},
		adsCache: adscache.ADSCache{
			Listeners: make(map[string]resources.Listener),
			Clusters:  make(map[string]resources.Cluster),
			Secrets:   make(map[string]resources.Secret),
		},
	}
}

// newSnapshotVersion increments the current snapshotVersion
// and returns as a string.
func (p *Processor) newSnapshotVersion() string {

	// Reset the snapshotVersion if it ever hits max size.
	if p.snapshotVersion == math.MaxInt64 {
		p.snapshotVersion = 0
	}

	// Increment the snapshot version & return as string.
	p.snapshotVersion++
	return strconv.FormatInt(p.snapshotVersion, 10)
}

// ProcessFile takes a file and generates an xDS snapshot
func (p *Processor) ProcessFile(file watcher.NotifyMessage) {

	// Parse file into object
	envoyConfig, err := parseYaml(file.FilePath)
	if err != nil {
		p.Errorf("error parsing yaml file: %+v", err)
		return
	}

	// Parse Secrets
	for _, s := range envoyConfig.Secrets {
		key, err := os.ReadFile(s.PrivateKeyFile)
		if err != nil {
			continue
		}
		ca, err := os.ReadFile(s.CertificateChainFile)
		if err != nil {
			continue
		}
		p.adsCache.AddSecret(s.Name, key, ca)
	}

	// Parse Listeners
	for _, l := range envoyConfig.Listeners {
		if len(l.SNIs) > 0 {
			cacheSecrets := p.adsCache.Secrets
			var snis []resources.SNI
			for _, sniCfg := range l.SNIs {
				var secretNames []string
				for _, secretName := range sniCfg.SecretNames {
					if _, ok := cacheSecrets[secretName]; !ok {
						continue
					}
					secretNames = append(secretNames, secretName)
				}
				if len(secretNames) == 0 {
					continue
				}

				var routes []resources.Route
				for _, r := range sniCfg.Routes {
					routes = append(routes, resources.Route{
						Name:    r.Name,
						Prefix:  r.Prefix,
						Cluster: r.ClusterNames[0],
					})
				}

				sni := resources.SNI{
					ServerNames: sniCfg.ServerNames,
					SecretNames: secretNames,
					Routes:      routes,
				}
				snis = append(snis, sni)
			}

			p.adsCache.AddListener(l.Name, l.Address, l.Port, snis)
		}
	}

	// Parse Clusters
	for _, c := range envoyConfig.Clusters {
		p.adsCache.AddCluster(c.Name)

		// Parse endpoints
		for _, e := range c.Endpoints {
			p.adsCache.AddEndpoint(c.Name, e.Address, e.Port)
		}
	}

	// Create the snapshot that we'll serve to Envoy
	snapshot := cache.NewSnapshot(
		p.newSnapshotVersion(),        // version
		[]types.Resource{},            // endpoints
		p.adsCache.ClusterContents(),  // clusters
		[]types.Resource{},            // routes
		p.adsCache.ListenerContents(), // listeners
		[]types.Resource{},            // runtimes
		p.adsCache.SecretContents(),   // secrets
	)

	if err := snapshot.Consistent(); err != nil {
		p.Errorf("snapshot inconsistency: %+v\n\n\n%+v", snapshot, err)
		return
	}
	p.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := p.cache.SetSnapshot(p.nodeID, snapshot); err != nil {
		p.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}
}
