package config

type Config interface {
	Parse(path string) *AppConfig
}

type HttpListen struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Log struct {
	LogFile  string `yaml:"file_path"`
	LogLevel string `yaml:"level"`
}

type AppConfig struct {
	HttpListen `yaml:"http_server"`
	Log        `yaml:"log"`
}
