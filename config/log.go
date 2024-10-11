package config

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	Dir        string `mapstructure:"dir" json:"dir" yaml:"dir"`
	FileName   string `mapstructure:"file_name" json:"file_name" yaml:"file_name"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // days
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
