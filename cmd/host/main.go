package main

import (
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
	h := config.New()
	h.InitFile(f)

	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}

	h.SetValue("key4", "val4")

	pluginCmd, pluginCfg := plugin.PluginCmd()
	rootCmd.AddCommand(pluginCmd)

	h.SetPluginCfg("pluginA", pluginCfg)

	h.MergePluginCfg()

	err = h.Save()
	if err != nil {
		log.Fatal(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
