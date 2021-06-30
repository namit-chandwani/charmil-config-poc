package plugin

import (
	"fmt"

	"github.com/namit-chandwani/charmil-config-poc/core/config"
	"github.com/spf13/cobra"
)

func PluginCmd() (*cobra.Command, map[string]interface{}) {
	config.New()

	config.SetValue("key5", "val5")
	config.SetValue("key6", "val6")
	config.SetValue("key7", "val7")
	config.SetValue("key8", "val8")

	fmt.Println("Plugin config map: ", config.GetAllSettings())

	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Plugin root command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Plugin root command called")
		},
	}

	return cmd, config.GetAllSettings()
}
