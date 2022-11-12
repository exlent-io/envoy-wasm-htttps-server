package processor

import (
	"fmt"
	"github.com/exlent-io/envoy-wasm-htttps-server/apis/v1alpha1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseYamlContents(t *testing.T) {
	testCases := []struct {
		name    string
		content []byte
		want    *v1alpha1.EnvoyConfig
		wantErr error
	}{
		{
			name: "sni",
			content: []byte(`
name: sni_config
spec:
  listeners:
  - name: listener_0
    address: 0.0.0.0
    port: 9000
    sni:
      - server_names:
          - example.com
          - www.example.com
      - secret_names:
          - example_com
      - routes:
        - name: echo_route
          prefix: /
          clusters:
            - echo
  clusters:
  - name: echo
    endpoints:
    - address: 127.0.0.1
      port: 9101
    - address: 127.0.0.1
      port: 9102
  secrets:
    - name: example_com
      private_key_file: "example_com_key.pem"
      certificate_chain_file: "example_com_cert.pem"
`),
			want: &v1alpha1.EnvoyConfig{
				Name: "sni_config",
				Spec: v1alpha1.Spec{
					Listeners: []v1alpha1.Listener{
						{
							Name:    "listener_0",
							Address: "0.0.0.0",
							Port:    uint32(9000),
							SNIs: []v1alpha1.SNI{
								{
									ServerNames: []string{
										"example.com",
										"www.example.com",
									},
									SecretNames: []string{
										"example_com",
									},
									Routes: []v1alpha1.Route{
										{
											Name:   "echo_route",
											Prefix: "/",
											ClusterNames: []string{
												"echo",
											},
										},
									},
								},
							},
						},
					},
					Clusters: []v1alpha1.Cluster{
						{
							Name: "echo",
							Endpoints: []v1alpha1.Endpoint{
								{
									Address: "127.0.0.1",
									Port:    uint32(9101),
								},
								{
									Address: "127.0.0.1",
									Port:    uint32(9102),
								},
							},
						},
					},
					Secrets: []v1alpha1.Secret{
						{
							Name:                 "example_com",
							PrivateKeyFile:       "example_com_key.pem",
							CertificateChainFile: "example_com_cert.pem",
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseYamlContents(tc.content)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				assert.Equal(t, tc.want, got)
			}
			fmt.Println(got)
		})
	}
}
