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
