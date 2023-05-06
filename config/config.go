package config

type MySqlConfig struct {
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Database string `mapstructure:"database"`
}

type ServerConfig struct {
	Host  string      `mapstructure:"host"`
	Port  int         `mapstructure:"port"`
	MySql MySqlConfig `mapstructure:"mysql"`
}
