package config

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int16  `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}
