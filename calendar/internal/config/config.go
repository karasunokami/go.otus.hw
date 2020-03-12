package config

type Config interface {
	Parse(path string) *AppConfig
}

type HttpListen struct {
	Ip   string
	Port string
}

type LogLevel string

//const (
//	Info  LogLevel = "info"
//	Warn  LogLevel = "warn"
//	Error LogLevel = "error"
//	Debug LogLevel = "debug"
//)

type AppConfig struct {
	HttpListen `yaml:"http_server"`
	LogFile    string   `yaml:"log_file"`
	LogLevel   LogLevel `yaml:"log_level"`
}
