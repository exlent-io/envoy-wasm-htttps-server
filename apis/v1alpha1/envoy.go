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

package v1alpha1

type EnvoyConfig struct {
	Name string `yaml:"name"`
	Spec `yaml:"spec"`
}

type Spec struct {
	Listeners []Listener `yaml:"listeners"`
	Clusters  []Cluster  `yaml:"clusters"`
	Secrets   []Secret   `yaml:"secrets"`
}

type Listener struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    uint32 `yaml:"port"`
	SNIs    []SNI  `yaml:"sni"`
}

type Route struct {
	Name         string   `yaml:"name"`
	Prefix       string   `yaml:"prefix"`
	ClusterNames []string `yaml:"clusters"`
}

type Cluster struct {
	Name      string     `yaml:"name"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Address string `yaml:"address"`
	Port    uint32 `yaml:"port"`
}

// SNI server name indication
type SNI struct {
	ServerNames []string `yaml:"server_names"`
	SecretNames []string `yaml:"secret_names"`
	Routes      []Route  `yaml:"routes"`
}

type Secret struct {
	Name                 string `yaml:"name"`
	PrivateKeyFile       string `yaml:"private_key_file"`
	CertificateChainFile string `yaml:"certificate_chain_file"`
}
