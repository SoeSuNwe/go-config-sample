package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	AppConfig struct {
		DatabaseUrl string `yaml:"database_url"`
		Host        string `yaml:"host"`
	}

	ConfigFile map[string]*AppConfig
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host          string        `json:"host"`
	Port          string        `json:"port"`
	KafkaSettings KafkaSettings `json:"kafka_settings"`
}

type KafkaSettings struct {
	BootstrapServers struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"bootstrap_servers"`
	Topics string `json:"topics"`
}

func LoadConfig(env string) (*AppConfig, error) {
	configFile := ConfigFile{}
	file, _ := os.Open("config.yml")
	defer file.Close()
	decoder := yaml.NewDecoder(file)

	// Always check for errors!
	if err := decoder.Decode(&configFile); err != nil {
		return nil, err
	}

	appConfig, ok := configFile[env]
	if !ok {
		return nil, fmt.Errorf("no such environment: %s", env)
	}

	return appConfig, nil
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}

func main() {
	// appConfig, err := LoadConfig(os.Args[1])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("config: %+v\n", appConfig)
	// fmt.Printf("config database: %+v\n", appConfig.DatabaseUrl)
	// fmt.Printf("config database: %+v\n", appConfig.Host)

	fmt.Println("**************************************************************")
	conf, err := LoadConfiguration("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("congig", conf.KafkaSettings.BootstrapServers)
	fmt.Println("congig", conf.KafkaSettings.Topics)

}
