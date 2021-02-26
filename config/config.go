package config

//based on: https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const defaultConfigPath = "../user-store/config.yml"

// Config struct for service config
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Host      string        `yaml:"host"`
		Port      int           `yaml:"port"`
		User      string        `yaml:"user"`
		Password  string        `yaml:"password"`
		Timeout   time.Duration `yaml:"timeout"`
		Name      string        `yaml:"name"`
		Schema    string        `yaml:"schema"`
		Migration struct {
			Run   bool   `yaml:"run"`
			Files string `yaml:"files"`
		} `yaml:"migration"`
	} `yaml:"db"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", defaultConfigPath, "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}
