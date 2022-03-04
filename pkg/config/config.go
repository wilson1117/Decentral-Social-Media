package config

import (
	"bytes"
	"os"

	"github.com/spf13/viper"
)

type NodeConfig struct {
	v *viper.Viper
}

func New(filepath string) (*NodeConfig, error) {
	var (
		configFile []byte
		err        error
	)

	if configFile, err = os.ReadFile(filepath); err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigType("json")

	if err = v.ReadConfig(bytes.NewBuffer(configFile)); err != nil {
		return nil, err
	}

	return &NodeConfig{v: v}, nil
}

func (nc NodeConfig) GetConfig(part string, cfg *map[string]interface{}) {
	sub := nc.v.Sub(part)

	for k, v := range *cfg {
		sub.SetDefault(k, v)
		switch v.(type) {
		case string:
			(*cfg)[k] = sub.GetString(k)
		case int:
			(*cfg)[k] = sub.GetInt(k)
		case []string:
			(*cfg)[k] = sub.GetStringSlice(k)
		case bool:
			(*cfg)[k] = sub.GetBool(k)
		}
	}
}
