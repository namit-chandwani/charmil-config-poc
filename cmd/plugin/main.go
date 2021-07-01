package plugin

import (
	"fmt"

	"github.com/namit-chandwani/charmil-config-poc/core/config"
	"github.com/spf13/cobra"
)

func PluginCmd() (*cobra.Command, map[string]interface{}) {
	h := config.New()

	h.SetValue("key5", "val5")
	h.SetValue("key6", "val6")
	h.SetValue("key7", "val7")
	h.SetValue("key8", "val8")

	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Plugin root command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Plugin root command called")
		},
	}

	return cmd, h.GetAllSettings()
}
