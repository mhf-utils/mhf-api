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
	Launcher Launcher
	Locales  []string
	Logger   Logger
	Mhfdat   Mhfdat
	NewRelic NewRelic
}

type Info struct {
	FilePath string // FilePath
	Enable   bool   // To enable or disable the router linked
}

type Launcher struct {
	En Info // LauncherInfo for En version
	Fr Info // LauncherInfo for Fr version
	Jp Info // LauncherInfo for Jp version
}

type Mhfdat struct {
	En Info // MhfdatInfo for En version
	Fr Info // MhfdatInfo for Fr version
	Jp Info // MhfdatInfo for Jp version
}

type Logger struct {
	Format   string
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
		{Name: "launcher"},
		{Name: "locales"},
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
