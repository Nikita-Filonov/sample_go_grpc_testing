package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func getEnv() (Env, error) {
	env := os.Getenv("ENV")

	if len(env) == 0 {
		return "", fmt.Errorf("cannot parse env variable")
	}

	return Env(env), nil
}

func readConfig(configPath string, config interface{}) error {
	if configPath == `` {
		return fmt.Errorf(`no config path`)
	}

	configBytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(configBytes, config); err != nil {
		return err
	}

	return nil
}

func NewConfig() (Config, error) {
	var config Config

	env, err := getEnv()

	if err != nil {
		return Config{}, err
	}

	err = readConfig(fmt.Sprintf("../infrastructure/config-%s.yml", env), &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
