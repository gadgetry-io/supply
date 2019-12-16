package common

import (
	"github.com/spf13/viper"
)

// GetUserConfig - Common Utility Function to Get User Config Value from $HOME/.supply/config.yaml
func GetUserConfig(mapName string, mapKey string) string {
	DebugMessage("[common] func GetUserConfig()")
	configMap := viper.GetStringMapString(mapName)
	configValue := configMap[mapKey]
	return configValue
}
