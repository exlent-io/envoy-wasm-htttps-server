package adscache

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/exlent-io/envoy-wasm-htttps-server/internal/resources"
)

type ADSCache struct {
	Listeners map[string]resources.Listener
	Clusters  map[string]resources.Cluster
	Secrets   map[string]resources.Secret
}

func (c *ADSCache) ClusterContents() []types.Resource {
	var r []types.Resource

	for _, cluster := range c.Clusters {
		r = append(r, resources.MakeCluster(cluster.Name, cluster.Endpoints))
	}

	return r
}

func (c *ADSCache) ListenerContents() []types.Resource {
	var r []types.Resource

	for _, listener := range c.Listeners {
		r = append(r, resources.MakeListener(listener.Name, listener.Address, listener.Port, nil, listener.SNIs))
	}

	return r
}

func (c *ADSCache) SecretContents() []types.Resource {
	var r []types.Resource

	for _, secret := range c.Secrets {
		r = append(r, resources.MakeSecret(secret.Name, secret.PrivateKey, secret.CertificateChain))
	}

	return r
}

func (c *ADSCache) AddSecret(name string, privateKey, certificateChain []byte) {
	c.Secrets[name] = resources.Secret{
		Name:             name,
		PrivateKey:       privateKey,
		CertificateChain: certificateChain,
	}
}

func (c *ADSCache) AddListener(name string, address string, port uint32, snis []resources.SNI) {
	c.Listeners[name] = resources.Listener{
		Name:    name,
		Address: address,
		Port:    port,
		SNIs:    snis,
	}
}

func (c *ADSCache) AddCluster(name string) {
	c.Clusters[name] = resources.Cluster{
		Name: name,
	}
}

func (c *ADSCache) AddEndpoint(clusterName, upstreamHost string, upstreamPort uint32) {
	cluster := c.Clusters[clusterName]

	cluster.Endpoints = append(cluster.Endpoints, resources.Endpoint{
		UpstreamHost: upstreamHost,
		UpstreamPort: upstreamPort,
	})

	c.Clusters[clusterName] = cluster
}
