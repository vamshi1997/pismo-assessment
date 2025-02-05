package boot

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	cfg Config
)

type Config struct {
	AppConfig App `mapstructure:"app"`
}

type App struct {
	Server struct {
		Host string
		Port int
	} `mapstructure:"server"`
	DB struct {
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		Port     int    `mapstructure:"port"`
		Charset  string `mapstructure:"charset"`
	} `mapstructure:"db"`
}

func InitConfig() {
	viper.SetConfigName("/configs/default") // name of config file (without extension)
	viper.SetConfigType("toml")             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/app")             // path to look for the config file in
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Configs are loaded successfully ...")

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Error unmarshaling config: %s\n", err)
		return
	}

	fmt.Println("Configs are mapped properly ...")
}

func GetConfig() Config {
	return cfg
}
