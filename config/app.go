package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Url     string `mapstructure:"url" json:"url" yaml:"url"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
}
