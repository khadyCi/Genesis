package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
}

func LoadConfig(cfg *Config, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	fmt.Println("Configuración cargada con éxito")
	return nil
}
