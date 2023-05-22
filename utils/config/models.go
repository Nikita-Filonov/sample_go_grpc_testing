package config

type Env string

const (
	DevEnv    Env = "dev"
	LocalEnv  Env = "local"
	StableEnv Env = "stable"
)

type Logger struct {
	Level     string `yaml:"level"`
	IsDevMode bool   `yaml:"isDevMode"`
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
