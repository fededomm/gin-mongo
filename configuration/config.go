package configuration

type DbConfig struct {
	ConnectionString string `mapstructure:connectionstring`
}

type RouterConf struct {
	Router string `mapstructure:router`
}
