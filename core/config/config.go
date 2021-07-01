package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// File contains the fields used to point to a local config file
type File struct {
	Name string
	Type string
	Path string
}

// Handler represents a wrapper around Viper
type Handler struct {
	vp *viper.Viper
}

// Plugin represents the config values map imported from a plugin
type Plugin map[string]interface{}

var pluginCfg = make(map[string]Plugin)

// New returns a new instance of the handler
func New() *Handler {
	h := &Handler{vp: viper.New()}
	return h
}

// InitFile selects a local config file based on the specified file configuration
func (h *Handler) InitFile(f File) {
	h.vp.SetConfigName(f.Name) // name of config file (without extension)
	h.vp.SetConfigType(f.Type) // REQUIRED if the config file does not have the extension in the name
	h.vp.AddConfigPath(f.Path) // path to look for the config file in
}

// Load imports config values from the local config file
// which was initialized while calling the InitFile method
func (h *Handler) Load() error {
	err := h.vp.ReadInConfig() // Find and read the config file
	return err
}

// Save writes the current config into the local config file
// which was initialized while calling the InitFile method
func (h *Handler) Save() error {
	err := h.vp.WriteConfig()
	return err
}

// SetValue stores the specified key, value pair in the current config
func (h *Handler) SetValue(key string, value interface{}) {
	h.vp.Set(key, value)
}

// GetValue returns the value of the key passed in as argument from the current config
func (h *Handler) GetValue(key string) (interface{}, error) {
	if h.vp.IsSet(key) {
		return h.vp.Get(key), nil
	}
	return nil, fmt.Errorf("Key doesn't exist")
}

// GetAllSettings returns the current config in the form of a map
func (h *Handler) GetAllSettings() map[string]interface{} {
	return h.vp.AllSettings()
}

// SetPluginCfg stores the imported plugin config as a key, value pair in a map.
// Here, the key represents name of the plugin and the value being its config map.
// For eg. Key: "pluginA", Value: map[key5:value5 key6:value6 key7:value7 key8:value8]
func (h *Handler) SetPluginCfg(pluginName string, p Plugin) {
	pluginCfg[pluginName] = p
}

// MergePluginCfg stores the pluginCfg map into the current config
func (h *Handler) MergePluginCfg() {
	h.SetValue("plugins", pluginCfg)
}
