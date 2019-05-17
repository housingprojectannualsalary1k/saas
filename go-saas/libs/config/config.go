package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Database struct {
		Server        string
		Port          int
		Dbname		  string
		Username	  string
		Password	  string
		Charset		  string
	}

	Session struct {
		Server        string
		Port          int
		Password	  string
	}
}
var config Config

func Get() *Config{
	return &config
}

func GetFileAndLoad(filePath string) *Config{
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(fmt.Sprintf("account slave db connect errr: %s", err))
	}

	return &config
}