package cli

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path                             string
	Username, Token, App, Collection string
}

func (c *Config) Save() error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(c.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
