package config

import (
	"encoding/json"
	"errors"
	"os"
)

// Config ..
type Config struct {
	Version                 string `json:"version,omitempty"`
	ObsExePath              string `json:"obsExePath,omitempty"`
	ObsSettingsPath         string `json:"obsSettingsPath,omitempty"`
	MaxRecordingSeconds     int    `json:"maxRecordingSeconds,omitempty"`
	MaxHoursRetention       int    `json:"maxHoursRetention,omitempty"`
	ObsWebsocketHost        string `json:"obsWebsocketHost,omitempty"`
	ObsWebsocketPort        int    `json:"obsWebsocketPort,omitempty"`
	ObsWebsocketWaitSeconds int    `json:"obsWebsocketWaitSeconds,omitempty"`
	IsReady                 bool   `json:"-"`
}

// AppConfig - The current app config populated after ReadConfig
var AppConfig Config

// ConfigVersion - The version of config json format
var ConfigVersion string = "1.0.0"

// ReadConfig - Read the config.json file and return a config struct
func ReadConfig() (Config, error) {
	file, err := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	retval := Config{}
	decoder.Decode(&retval)
	if err != nil {
		return retval, err
	}
	if retval.Version == ConfigVersion {
		retval.IsReady = true
		AppConfig = retval
	} else {
		return AppConfig, errors.New("config version invalid")
	}
	return AppConfig, nil
}
