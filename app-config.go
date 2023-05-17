package main

import (
	_ "embed"
	"fmt"
	AppConf "gin-mongo/configuration"
	"os"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// Configuration struct
type Config struct {
	Log LogConfig `yaml:"log"`
	App AppConfig `yaml:"config"`
}

type LogConfig struct {
	Level      int  `yaml:"level"`
	EnableJSON bool `yaml:"enablejson"`
}

type AppConfig struct {
	Database      AppConf.DbConfig      `yaml:"database" mapstructure:"database" json:"database"`
	Router        AppConf.RouterConf    `yaml:"router" mapstructure:"router" json:"router"`
	Observability AppConf.Observability `yaml:"observability" mapstructure:"observability" json:"observability"`
	ServiceName   string                `yaml:"servicename" mapstructure:"servicename" json:"servicename"`
}

var DefaultConfig = Config{
	Log: LogConfig{
		Level:      -1,
		EnableJSON: false,
	},
	App: AppConfig{
		Database: AppConf.DbConfig{
			ConnectionString: "mongodb://root:example127.0.0.2:27017",
		},
		Router: AppConf.RouterConf{
			Router: "0.0.0.0:8085",
		},
		Observability: AppConf.Observability{
			ServiceName: "gin-mongo",
			Endpoint:    "192.168.3.109:4317",
			Enable:      false,
		},
	},
}

// Default config file.
//
//go:embed configuration/config.yaml
var projectConfigFile []byte

const ConfigFileEnvVar = "GIN_MONGO_FILE_PATH"
const ConfigurationName = "gin-mongo"

func ReadConfig() (*Config, error) {

	configPath := os.Getenv(ConfigFileEnvVar)
	var cfgContent []byte
	var err error
	if configPath != "" {
		if _, err := os.Stat(configPath); err == nil {
			log.Info().Str("cfg-file-name", configPath).Msg("reading config")
			cfgContent, err = util.ReadFileAndResolveEnvVars(configPath)
			log.Info().Msg("++++CFG:" + string(cfgContent))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("the %s env variable has been set but no file cannot be found at %s", ConfigFileEnvVar, configPath)
		}
	} else {
		log.Warn().Msgf("The config path variable %s has not been set. Reverting to bundled configuration", ConfigFileEnvVar)
		cfgContent = util.ResolveConfigValueToByteArray(projectConfigFile)
		// return nil, fmt.Errorf("the config path variable %s has not been set; please set", ConfigFileEnvVar)
	}

	appCfg := DefaultConfig
	err = yaml.Unmarshal(cfgContent, &appCfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if !appCfg.Log.EnableJSON {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	zerolog.SetGlobalLevel(zerolog.Level(appCfg.Log.Level))

	return &appCfg, nil
}
