package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type File struct {
	Name string
	Type string
	Path string
}

type Plugin map[string]interface{}

var (
	v         *viper.Viper
	pluginCfg = make(map[string]Plugin)
)

func New() {
	v = viper.New()
}

func InitFile(f File) {
	v.SetConfigName(f.Name) // name of config file (without extension)
	v.SetConfigType(f.Type) // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(f.Path) // path to look for the config file in
}

func Load() error {
	err := v.ReadInConfig() // Find and read the config file
	return err
}

func Save() error {
	// writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	err := v.WriteConfig()
	return err
}

func SetValue(key string, value interface{}) {
	v.Set(key, value)
}

func GetValue(key string) (interface{}, error) {
	if v.IsSet(key) {
		return v.Get(key), nil
	}
	return nil, fmt.Errorf("Key doesn't exist")
}

func GetAllSettings() map[string]interface{} {
	return v.AllSettings()
}

func SetPluginCfg(pluginName string, p Plugin) {
	pluginCfg[pluginName] = p
}

func MergePluginCfg() {
	SetValue("plugins", pluginCfg)
}
