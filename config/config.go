package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig reads configuration from a TOML file
func InitConfig(fileName string, additionalDirs []string) (string, error) {
	viper.SetConfigName(fileName)

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	for _, dir := range additionalDirs {
		viper.AddConfigPath(dir)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	configFile := viper.ConfigFileUsed()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
	})

	return configFile, nil
}

// GetConfigString returns a string value from configuration file
func GetConfigString(key string) string {
	return viper.GetString(key)
}

// GetConfigBool returns a bool value from configuration file
func GetConfigBool(key string) bool {
	return viper.GetBool(key)
}

// GetConfigInt returns an int value from configuration file
func GetConfigInt(key string) int {
	return viper.GetInt(key)
}

// GetConfigFloat64 returns a float64 value from configuration file
func GetConfigFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetConfigStringSlice returns a []string value from configuration file
func GetConfigStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetConfigDuration returns a time.Duration value from configuration file
func GetConfigDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

// GetConfigTime returns a time.Time value from configuration file
func GetConfigTime(key string) time.Time {
	return viper.GetTime(key)
}
