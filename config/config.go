package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Database contaisns database connection information.
type Database struct {
	Dialect    string
	DataSource string
}

// Configs is structure of configs.
// configを追加したら新たに構造体を定義して追加すれば使えるようになる
type Configs struct {
	Database Database `yaml:"database"`
}

var (
	conf *Configs
)

// GetConfigs 設定の取得
func GetConfigs() (*Configs, error) {
	c := &Configs{}
	var err error
	env := os.Getenv("ENV_GO")
	b, err := ioutil.ReadFile("./config/" + env + ".yml")
	if err != nil {
		return nil, err
	}	
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}