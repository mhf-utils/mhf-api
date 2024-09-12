package config

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	Logger   Logger
	Mhfdat   Mhfdat
	NewRelic NewRelic
}

type Logger struct {
	Format   string
	FilePath string
}

type Mhfdat struct {
	En MhfdatInfo
	Fr MhfdatInfo
	Jp MhfdatInfo
}

type MhfdatInfo struct {
	FilePath string
}
type NewRelic struct {
	License                 string
	AppName                 string
	AppLogForwardingEnabled bool
}

type ConfigFile struct {
	Name string
}

var GlobalConfig *Config

func init() {
	env := os.Getenv("ENVIRONMENT")
	fmt.Println("Config:init ENVIRONMENT=" + "'" + env + "'")
	var err error
	GlobalConfig, err = LoadConfig(env)
	if err != nil {
		preventClose(fmt.Sprintf("Failed to load config: %s", err.Error()))
	}
}

func LoadConfig(env string) (*Config, error) {
	var config Config

	config_files := []ConfigFile{
		{Name: "base"},
		{Name: "logger"},
		{Name: "mhfdat"},
		{Name: "newrelic"},
	}

	viper.SetConfigType("json")
	path := fmt.Sprintf("./config/%s", env)
	fmt.Println("Config:LoadConfig PATH=" + "'" + path + "'")
	viper.AddConfigPath(path)

	for _, config_file := range config_files {
		viper.SetConfigName(config_file.Name)

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
		if err := viper.Unmarshal(&config); err != nil {
			return nil, err
		}
	}

	if config.Host == "" {
		config.Host = getOutboundIP4().To4().String()
	}
	return &config, nil
}

func preventClose(text string) {
	fmt.Println("\nFailed to start mhf-api:\n" + text)
	go wait()
	fmt.Println("\nPress Enter/Return to exit...")
	fmt.Scanln()
	os.Exit(0)
}

func wait() {
	for {
		time.Sleep(time.Millisecond * 100)
	}
}

func getOutboundIP4() net.IP {
	conn, err := net.Dial("udp4", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.To4()
}
