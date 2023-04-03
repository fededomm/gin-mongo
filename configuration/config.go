package configuration

type DbConfig struct {
	ConnectionString string `mapstructure:connectionstring`
}

type RouterConf struct {
	Router string `mapstructure:router`
}

type Observability struct {
	ServiceName string `mapstructure:serviceName`
	Endpoint    string `mapstructure:endpoint`
	Enable      bool   `mapstructure:enable`
}
