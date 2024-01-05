package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Config YamlConfig

type YamlConfig struct {
	App      App            `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

type App struct {
	Name         string `yaml:"name"`
	Version      string `yaml:"version"`
	Introduction string `yaml:"introduction"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// InitConfigDev 初始化yaml config
func InitConfigDev() error {
	file, err := os.ReadFile("./config-dev.yaml")
	if err != nil {
		LOG.Println("read config-dev.yaml fail", err)
		return err
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		LOG.Println("yaml unmarshal fail")
		return err
	}
	return nil
}

// PrintConfig 打印配置
func PrintConfig() {
	fmt.Println("**************************************************")
	fmt.Println("*                        ")
	fmt.Println("*   Welcome to code-go   ")
	fmt.Printf("*   Introduction: %s\n", Config.App.Introduction)
	fmt.Printf("*   Click url：http://localhost%s\n", Config.Server.Port)
	fmt.Println("*                        ")
	fmt.Println("**************************************************")
}
