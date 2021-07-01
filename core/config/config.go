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

type Handler struct {
	vp *viper.Viper
}

var pluginCfg = make(map[string]Plugin)

func New() *Handler {
	h := &Handler{vp: viper.New()}
	return h
}

func (h *Handler) InitFile(f File) {
	h.vp.SetConfigName(f.Name) // name of config file (without extension)
	h.vp.SetConfigType(f.Type) // REQUIRED if the config file does not have the extension in the name
	h.vp.AddConfigPath(f.Path) // path to look for the config file in
}

func (h *Handler) Load() error {
	err := h.vp.ReadInConfig() // Find and read the config file
	return err
}

func (h *Handler) Save() error {
	// writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	err := h.vp.WriteConfig()
	return err
}

func (h *Handler) SetValue(key string, value interface{}) {
	h.vp.Set(key, value)
}

func (h *Handler) GetValue(key string) (interface{}, error) {
	if h.vp.IsSet(key) {
		return h.vp.Get(key), nil
	}
	return nil, fmt.Errorf("Key doesn't exist")
}

func (h *Handler) GetAllSettings() map[string]interface{} {
	return h.vp.AllSettings()
}

func (h *Handler) SetPluginCfg(pluginName string, p Plugin) {
	pluginCfg[pluginName] = p
}

func (h *Handler) MergePluginCfg() {
	h.SetValue("plugins", pluginCfg)
}
