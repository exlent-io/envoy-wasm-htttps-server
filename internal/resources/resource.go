// Copyright 2020 Envoyproxy Authors
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

package resources

import (
	"log"
	"time"

	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"

	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/ptypes"
)

func MakeCluster(name string, eps []Endpoint) *cluster.Cluster {
	var endpoints = make([]*endpoint.LbEndpoint, 0, len(eps))

	for _, e := range eps {
		endpoints = append(endpoints, &endpoint.LbEndpoint{
			HostIdentifier: &endpoint.LbEndpoint_Endpoint{
				Endpoint: &endpoint.Endpoint{
					Address: &core.Address{
						Address: &core.Address_SocketAddress{
							SocketAddress: &core.SocketAddress{
								Protocol: core.SocketAddress_TCP,
								Address:  e.UpstreamHost,
								PortSpecifier: &core.SocketAddress_PortValue{
									PortValue: e.UpstreamPort,
								},
							},
						},
					},
				},
			},
		})
	}

	return &cluster.Cluster{
		Name:                 name,
		ConnectTimeout:       ptypes.DurationProto(2 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_STRICT_DNS},
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment: &endpoint.ClusterLoadAssignment{
			ClusterName: name,
			Endpoints: []*endpoint.LocalityLbEndpoints{{
				LbEndpoints: endpoints,
			}},
		},
	}
}

func MakeListener(name, address string, port uint32, rts []Route, snis []SNI) *listener.Listener {
	var virtualHostName = "local_service"

	filterChains := make([]*listener.FilterChain, 0, len(snis))
	for _, sni := range snis {
		sdsTls := makeDownstreamTlsContext(sni.SecretNames)
		sdsCfg, err := ptypes.MarshalAny(sdsTls)
		if err != nil {
			panic(err)
		}
		filters := make([]*listener.Filter, 0, len(sni.Routes))

		rte := &route.RouteConfiguration{
			Name: name,
			VirtualHosts: []*route.VirtualHost{{
				Name:    virtualHostName,
				Domains: []string{"*"},
				Routes:  makeRoutes(sni.Routes),
			}},
		}
		routerConfig, err := ptypes.MarshalAny(&router.Router{})
		if err != nil {
			log.Fatal(err)
		}
		manager := &hcm.HttpConnectionManager{
			CodecType:  hcm.HttpConnectionManager_AUTO,
			StatPrefix: "ingress_http",
			RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
				RouteConfig: rte,
			},
			HttpFilters: []*hcm.HttpFilter{{
				Name: wellknown.Router,
				ConfigType: &hcm.HttpFilter_TypedConfig{
					TypedConfig: routerConfig,
				},
			}},
		}
		pbst, err := ptypes.MarshalAny(manager)
		if err != nil {
			log.Fatal(err)
		}
		filters = append(filters, &listener.Filter{
			Name: wellknown.HTTPConnectionManager,
			ConfigType: &listener.Filter_TypedConfig{
				TypedConfig: pbst,
			},
		})
		filterChains = append(filterChains, &listener.FilterChain{
			FilterChainMatch: makeFilterChainMatch(sni.ServerNames),
			TransportSocket: &core.TransportSocket{
				Name: "envoy.transport_sockets.tls",
				ConfigType: &core.TransportSocket_TypedConfig{
					TypedConfig: sdsCfg,
				},
			},
			Filters: filters,
		})
	}

	return &listener.Listener{
		Name: name,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  address,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: filterChains,
	}
}

func MakeSecret(name string, key, crt []byte) *auth.Secret {
	return &auth.Secret{
		Name: name,
		Type: &auth.Secret_TlsCertificate{
			TlsCertificate: &auth.TlsCertificate{
				CertificateChain: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{InlineBytes: crt},
				},
				PrivateKey: &core.DataSource{
					Specifier: &core.DataSource_InlineBytes{InlineBytes: key},
				},
			},
		},
	}
}

func makeRoutes(routes []Route) []*route.Route {
	var rts []*route.Route
	for _, r := range routes {
		rts = append(rts, &route.Route{
			Match: &route.RouteMatch{
				PathSpecifier: &route.RouteMatch_Prefix{
					Prefix: r.Prefix,
				},
			},
			Action: &route.Route_Route{
				Route: &route.RouteAction{
					ClusterSpecifier: &route.RouteAction_Cluster{
						Cluster: r.Cluster,
					},
				},
			},
		})
	}
	return rts
}

func makeDownstreamTlsContext(secretNames []string) *auth.DownstreamTlsContext {
	sdsSecretConfigs := make([]*auth.SdsSecretConfig, 0, len(secretNames))
	for _, secretName := range secretNames {
		sdsSecretConfigs = append(sdsSecretConfigs, makeSdsSecretConfig(secretName))
	}
	return &auth.DownstreamTlsContext{
		CommonTlsContext: &auth.CommonTlsContext{
			TlsCertificateSdsSecretConfigs: sdsSecretConfigs,
		},
	}
}

func makeSdsSecretConfig(secretName string) *auth.SdsSecretConfig {
	return &auth.SdsSecretConfig{
		Name: secretName,
		SdsConfig: &core.ConfigSource{
			ConfigSourceSpecifier: &core.ConfigSource_Ads{
				Ads: &core.AggregatedConfigSource{},
			},
			ResourceApiVersion: core.ApiVersion_V3,
		},
	}
}

func makeFilterChainMatch(serverNames []string) *listener.FilterChainMatch {
	return &listener.FilterChainMatch{
		ServerNames: serverNames,
	}
}
