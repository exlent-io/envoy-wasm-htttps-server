//   Copyright Steve Sloka 2021
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package processor

import (
	"context"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/exlent-io/envoy-wasm-htttps-server/internal/resources"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"

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

	adsCache adscache.ADSCache

	configFilename string
}

func NewProcessor(cache cache.SnapshotCache, nodeID string, log logrus.FieldLogger, config string) *Processor {
	return &Processor{
		cache:           cache,
		nodeID:          nodeID,
		snapshotVersion: rand.Int63n(1000),
		FieldLogger:     log,
		adsCache: adscache.ADSCache{
			Listeners: make(map[string]resources.Listener),
			Clusters:  make(map[string]resources.Cluster),
			Secrets:   make(map[string]resources.Secret),
		},
		configFilename: config,
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
func (p *Processor) ProcessFile(ctx context.Context, file watcher.NotifyMessage) {

	// Parse file into object
	envoyConfig, err := parseYaml(p.configFilename)
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
	snapshot, err := cache.NewSnapshot(
		p.newSnapshotVersion(), // version
		map[resource.Type][]types.Resource{
			resource.ClusterType:  p.adsCache.ClusterContents(),
			resource.ListenerType: p.adsCache.ListenerContents(),
			resource.SecretType:   p.adsCache.SecretContents(),
			// resource.EndpointType: []types.Resource{},
			// resource.RouteType: []types.Resource{},
			// resource.RuntimeType: []types.Resource{},
		},
	)

	if err := snapshot.Consistent(); err != nil {
		p.Errorf("snapshot inconsistency: %+v\n\n\n%+v", snapshot, err)
		return
	}
	p.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := p.cache.SetSnapshot(ctx, p.nodeID, snapshot); err != nil {
		p.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}
}
