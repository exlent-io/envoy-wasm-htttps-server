package resources

type Listener struct {
	Name    string
	Address string
	Port    uint32
	SNIs    []SNI
}

type Route struct {
	Name    string
	Prefix  string
	Cluster string
}

type Cluster struct {
	Name      string
	Endpoints []Endpoint
}

type Endpoint struct {
	UpstreamHost string
	UpstreamPort uint32
}

type SNI struct {
	SecretNames []string
	ServerNames []string
	Routes      []Route
}

type Secret struct {
	Name             string
	PrivateKey       []byte
	CertificateChain []byte
}
