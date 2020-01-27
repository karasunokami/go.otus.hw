package config

type DB struct {
	Host 	 string `mapstructure:"host"`
	Port 	 string `mapstructure:"port"`
	User 	 string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
