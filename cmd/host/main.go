package main

import (
	"fmt"
	"log"

	"github.com/namit-chandwani/charmil-config-poc/cmd/plugin"
	"github.com/namit-chandwani/charmil-config-poc/core/config"
	"github.com/spf13/cobra"
)

var (
	f       config.File
	rootCmd *cobra.Command
)

func init() {
	f = config.File{
		Name: "config",
		Type: "yaml",
		Path: "./cmd/host",
	}

	rootCmd = &cobra.Command{
		Use:   "host",
		Short: "Host CLI",
	}
}

func main() {
	config.New()
	config.InitFile(f)

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.SetValue("key4", "val4")

	fmt.Println("Host config map: ", config.GetAllSettings())

	pluginCmd, pluginCfg := plugin.PluginCmd()
	rootCmd.AddCommand(pluginCmd)

	config.SetPluginCfg("pluginA", pluginCfg)

	fmt.Println("Host config map: ", config.GetAllSettings())

	config.MergePluginCfg()

	fmt.Println("Host config map: ", config.GetAllSettings())

	err = config.Save()
	if err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
