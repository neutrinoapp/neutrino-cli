package cmd

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/neutrinoapp/neutrino-cli/cli"
	"github.com/neutrinoapp/neutrino-client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var cfgFile string
var config *cli.Config

func ConfigPath() string {
	return cfgFile
}

func Config() *cli.Config {
	return config
}

func Client() *neutrino.ApiClient {
	return neutrino.NewApiClientCached(Config().App)
}

var RootCmd = &cobra.Command{
	Use:   "neutrino-cli",
	Short: "Neutrino cli",
	Long:  "",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.neutrino-cli.yaml)")
}

func initConfig() {
	if cfgFile == "" { // enable ability to specify config file via flag
		cfgFile = os.ExpandEnv("$HOME") + "/" + ".neutrino-cli.yaml"
	}

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		return
	}

	b, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	config = &cli.Config{}
	unmarshalError := yaml.Unmarshal(b, config)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
		return
	}

	config.Path = cfgFile

	neutrino.InitClient(cli.HTTP_ADDR, cli.WS_ADDR, config.Token, "cli")
}
