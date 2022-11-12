package processor

import (
	"fmt"
	"io/ioutil"

	"github.com/exlent-io/envoy-wasm-htttps-server/apis/v1alpha1"
	"gopkg.in/yaml.v2"
)

func parseYamlContents(bytes []byte) (*v1alpha1.EnvoyConfig, error) {
	var config v1alpha1.EnvoyConfig

	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// parseYaml takes in a yaml envoy config and returns a typed version
func parseYaml(file string) (*v1alpha1.EnvoyConfig, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading YAML file: %s\n", err)
	}

	return parseYamlContents(yamlFile)
}
