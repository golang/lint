package golintx

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl"
)

var configMap = map[string]*Config{}

type Config struct {
	Exclude     ExcludesConfig
	Initialisms []string
}

type ExcludesConfig struct {
	Categories []string
}

const golintxConfigFilename = ".golintx.hcl"

// http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func parseAndSetConfig(filename string) (*Config, error) {
	baseDir := filepath.Dir(filename)

	for dir := baseDir; ; dir = filepath.Dir(dir) {
		config, ok := configMap[dir]
		if ok {
			return config, nil
		}
		filename := filepath.Join(dir, golintxConfigFilename)
		exists, _ := fileExists(filename)
		if exists {
			config, err := parseConfig(filename)
			if err != nil {
				return nil, err
			}
			configMap[dir] = config
			return config, nil
		}
		if dir == filepath.Dir(dir) {
			break
		}
	}
	return nil, nil
}

func parseConfig(filename string) (*Config, error) {
	var config Config
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", filename, err)
	}

	obj, err := hcl.Parse(string(d))
	if err != nil {
		return nil, fmt.Errorf("Error parsing %s: %s", filename, err)
	}
	if err := hcl.DecodeObject(&config, obj); err != nil {
		return nil, err
	}
	return &config, nil
}

func excludeCategories(ps []Problem, cs []string) []Problem {
	newPs := make([]Problem, 0, len(ps))
	for _, p := range ps {
		if !strHas(cs, p.Category) {
			newPs = append(newPs, p)
		}
	}
	return newPs
}

func (config *Config) initialismMap() map[string]bool {
	if config == nil || len(config.Initialisms) == 0 {
		return commonInitialisms
	}
	m := make(map[string]bool, len(config.Initialisms))
	for _, i := range config.Initialisms {
		m[i] = true
	}
	return m
}

func strHas(ss []string, t string) bool {
	for _, s := range ss {
		if s == t {
			return true
		}
	}
	return false
}
