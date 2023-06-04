package config

type Env string
type LogLevel string

const (
	DebugLogLevel   LogLevel = "debug"
	InfoLogLevel    LogLevel = "info"
	WarningLogLevel LogLevel = "warning"
	ErrorLogLevel   LogLevel = "error"
)

type Logger struct {
	Level     LogLevel `yaml:"level"`
	IsDevMode bool     `yaml:"isDevMode"`
}

type GrpcService struct {
	Port               int32  `yaml:"port"`
	Host               string `yaml:"host"`
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
}

type Config struct {
	Logger          Logger      `yaml:"logger" validate:"required"`
	Articlesservice GrpcService `yaml:"articlesService" validate:"required"`
}
