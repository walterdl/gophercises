package urlshort

import (
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type site struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(fPath string, fallback http.Handler) (http.HandlerFunc, error) {
	rawYaml, err := readYamlFile(fPath)
	if err != nil {
		return nil, err
	}

	sites := make([]site, 0)
	err = yaml.Unmarshal(rawYaml, &sites)
	if err != nil {
		return nil, err
	}

	return MapHandler(toMap(sites), fallback), nil
}

func readYamlFile(fPath string) ([]byte, error) {
	f, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return result, err
}

func toMap(sites []site) map[string]string {
	result := make(map[string]string, len(sites))

	for _, s := range sites {
		result[s.Path] = s.Url
	}

	return result
}
