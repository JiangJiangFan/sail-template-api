package config

type Mysql struct {
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	User                string `mapstructure:"user" json:"user" yaml:"user"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdle             int    `mapstructure:"max_idle" json:"max_idle" yaml:"max_idle"`
	MaxOpen             int    `mapstructure:"max_open" json:"max_open" yaml:"max_open"`
	MaxLife             int    `mapstructure:"max_life" json:"max_life" yaml:"max_life"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
}
